package Lagrange

import (
	"Lagrange-SDK/errors"
	"Lagrange-SDK/events"
	"github.com/rotisserie/eris"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"

	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

type Core struct {
	ApiUrl                    *url.URL
	events                    map[string][]events.EventCallbackFunc
	lock                      sync.RWMutex
	err                       error
	Client                    *websocket.Conn
	Header                    http.Header
	handlePanic               func(any)
	retryCount, MaxRetryCount int
	apibase                   string
	BotQQ                     *int64
	groupQQ                   *int64
	token                     string

	done chan struct{}
}

type CoreOpt func(*Core)

func (c *Core) HandlePanic(h func(any)) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.handlePanic = h
}

func (c *Core) On(event events.EventName, callback events.EventCallbackFunc) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.events[string(event)] = append(c.events[string(event)], callback)
}

func (c *Core) ListenAndWait(ctx context.Context) (e error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	go func() {
		select {
		case <-interrupt:
			log.Info("用户关闭程序")
			cancel()
			if c.Client != nil {
				c.Client.Close()
			}
		case <-ctx.Done():
		}
	}()

	c.done = make(chan struct{}, 1)
	defer func() {
		log.Debug(e)
		if e != errors.ErrorContextCanceled {
			c.retryCount++
			if c.retryCount > c.MaxRetryCount {
				log.Info("超出重连次数")
				return
			}
			log.Warnf("连接出错，将进行第%d重连操作,按Ctrl+C取消重试", c.retryCount)
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(c.retryCount) * time.Second):
			}
			c.err = nil
			e = c.ListenAndWait(ctx)
			return
		}
	}()
	var err error
	c.Header = make(http.Header)
	c.Client, _, err = websocket.DefaultDialer.DialContext(ctx, "ws://"+c.ApiUrl.Host+"/ws", c.Header)
	if err != nil {
		return err
	}
	defer func() {
		if c.Client != nil {
			c.Client.Close()
		}
	}()
	c.retryCount = 0
	log.Info("连接成功到:" + c.ApiUrl.Host)
	go func() {
		defer close(c.done)
		for {
			_, message, err := c.Client.ReadMessage()
			select {
			case <-ctx.Done():
				c.err = errors.ErrorContextCanceled
				return
			default:
			}
			if err != nil {
				c.err = err
				return
			}
			log.Debug(string(message))
			event, _, ok, err := events.New(message)
			if err != nil {
				log.Error("error:", eris.ToString(err, true))
				continue
			}
			var (
				callbacks []events.EventCallbackFunc
			)
			c.lock.RLock()
			callbacks = c.events[ok]
			c.lock.RUnlock()
			go func() {
				defer func() {
					if err := recover(); err != nil {
						if c.handlePanic != nil {
							c.handlePanic(err)
						} else {
							log.Debugf("event handle function panic: %s \n%s", err, string(debug.Stack()))
						}
					}
				}()
				for _, v := range callbacks {
					v(ctx, event)
				}
			}()
		}
	}()

	<-c.done

	return c.err
}

func NewCore(api string, botQQ int64, opt ...CoreOpt) (*Core, error) {
	u, _ := url.Parse(api)
	c := &Core{
		ApiUrl:  u,
		apibase: api,
		events:  make(map[string][]events.EventCallbackFunc),
		lock:    sync.RWMutex{},
		done:    nil,
		//token:         accessToken,
		BotQQ:         &botQQ,
		MaxRetryCount: 10,
	}
	for _, o := range opt {
		o(c)
	}

	return c, nil
}
