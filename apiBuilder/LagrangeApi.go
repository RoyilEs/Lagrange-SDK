package apiBuilder

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/imroc/req/v3"
	"net/url"
)

type DoApi interface {
	Do(ctx context.Context) error
	DowithCallBack(ctx context.Context, callBack func(response *Response, err error)) error
	DoAndResponse(ctx context.Context) (*Response, error)
}

type Builder struct {
	url    string
	path   *string
	method *string
	action ApiName
	Params struct {
		GroupID  int64           `json:"group_id,omitempty"`
		UserID   int64           `json:"user_id,omitempty"`
		Message  []MessageStruct `json:"message,omitempty"`
		Duration int             `json:"duration,omitempty"`
	} `json:"params"`
}

func (b *Builder) BuildStringBody() (string, error) {
	body, err := json.Marshal(b)
	return string(body), err
}

func (b *Builder) Do(ctx context.Context) error {
	r, err := b.DoAndResponse(ctx)
	if err != nil {
		return err
	}
	if !r.Ok() {
		return fmt.Errorf(r.StatusMsg())
	}
	return nil
}

func (b *Builder) DowithCallBack(ctx context.Context, callBack func(response *Response, err error)) error {
	r, err := b.DoAndResponse(ctx)
	if err != nil {
		return err
	}
	if !r.Ok() {
		return fmt.Errorf(r.StatusMsg())
	}
	defer callBack(r, nil)
	return nil
}
func (b *Builder) DoAndResponse(ctx context.Context) (*Response, error) {
	body, err := b.BuildStringBody()
	if err != nil {
		return nil, err
	}
	log.Debug("requset", "body", body)
	client := req.SetContext(ctx)
	if b.path != nil {
		u, _ := url.JoinPath(b.url, *b.path)
		client.SetURL(u)
	} else {
		u, _ := url.JoinPath(b.url, string(b.action))
		client.SetURL(u)
	}

	if b.method != nil {
		client.Method = *b.method
	} else {
		client.Method = "POST"
	}
	resp := client.SetBodyString(body).Do()
	if resp.Err != nil {
		return nil, resp.Err
	}
	r := NewResponse(resp.Bytes())
	if !r.Ok() {
		return nil, fmt.Errorf(r.StatusMsg())
	}
	return r, nil
}

func New(url string) IMainFuncApi {
	return &Builder{url: url}
}
