package pixiv2

import (
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/global"
	"Lagrange-SDK/utils/new_req"
	"context"
	"encoding/json"
	"errors"
	"github.com/charmbracelet/log"
	"io/ioutil"
	"net/http"
	url1 "net/url"
	"strconv"
)

// 基响应
type BaseResp struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
	Message string `json:"message"`
	Wording string `json:"wording"`
	Echo    string `json:"echo,omitempty"`
}

// 基请求
type BaseReq struct {
}

type BaseInterface interface {
	GetReq() interface{}
	GetResp() interface{}
	Name() string
}

type GetPixivImageReq struct {
	BaseReq
}

type GetPixivImageData struct {
	Pid    int64    `json:"pid"`
	Page   int      `json:"page"`
	Uid    int64    `json:"uid"`
	Title  string   `json:"title"`
	User   string   `json:"user"`
	R18    int      `json:"r18"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Tags   []string `json:"tags"`
	Url    string   `json:"url"`
}

type GetPixivImageResp []GetPixivImageData

type GetPixivImageStruct struct {
	Req  *GetPixivImageReq
	Resp *GetPixivImageResp
}

func (g GetPixivImageStruct) Name() string {
	return PixivUrl2
}

func (g GetPixivImageStruct) GetReq() interface{} {
	return g.Req
}

func (g GetPixivImageStruct) GetResp() interface{} {
	return g.Resp
}

func GetPixivImage(client *http.Client, messageReq *GetPixivImageReq, form url1.Values) (err error, messageResp *GetPixivImageResp) {
	messageResp = new(GetPixivImageResp)
	return BaseServiceWithForm(client, GetPixivImageStruct{Req: messageReq, Resp: messageResp}, form), messageResp
}
func BaseServiceWithForm(client *http.Client, ReqResp BaseInterface, form url1.Values) (err error) {

	//zaplog.Logger.Info("messageReq:", messageReq)

	err, req := new_req.NewReqWithForm(ReqResp.Name(), form)

	if err != nil {
		//zaplog.Logger.Error(err.Error())
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		//zaplog.Logger.Error(err.Error())
		return err
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//zaplog.Logger.Error(err.Error())
		return err
	}

	err = json.Unmarshal(respData, ReqResp.GetResp())
	if err != nil {
		//zaplog.Logger.Error(err.Error())
		return err
	}
	//zaplog.Logger.Infof("%#v", messageResp)
	return nil
}

func GetPixivPidTitleUrl(client *http.Client, groupid int64, keyword string, r18 int, ctx context.Context) (pid int64, title string, url string, err error) {
	formData := url1.Values{}
	formData.Set("keyword", keyword)
	formData.Set("r18", strconv.Itoa(r18))
	err, resp := GetPixivImage(client, &GetPixivImageReq{}, formData)
	if err != nil {
		return 0, "", "", err
	}
	if len(*resp) == 0 {
		err = errors.New(global.ErrCmdPixTagNotFound + keyword)
		apiBuilder.New(global.BotUrl).SendGroupMsg(groupid).TextMsg(global.ErrCmdPixTagNotFound + keyword).Do(ctx)
		return 0, "", "", err
	}
	log.Debugf("resp: %v", (*resp)[0])
	return (*resp)[0].Pid, (*resp)[0].Title, (*resp)[0].Url, nil
}
