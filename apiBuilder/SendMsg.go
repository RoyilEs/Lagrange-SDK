package apiBuilder

import (
	"fmt"
)

type ISendReply interface {
	SendGroupMsg() ISendGroupMsg
	SendPrivateMsg() ISendPrivateMsg
}

type ISendGroupMsg interface {
	IMsg
	ISendReply
	ToGroupID(group int64) IMsg
}

type ISendPrivateMsg interface {
	IMsg
	ISendReply
	ToPrivateID(user int64) IMsg
}

type IMsg interface {
	TextMsg(text string) IMsg
	JsonMsg(json string) IMsg
	ImgMsg(img string) IMsg
	ImgBase64Msg(imgBase64 string) IMsg
	Face(ID int) IMsg
	DoApi
}

func (r *Request) SendReply(msgID int64) ISendReply {
	r.Params.Message += fmt.Sprintf("[CQ:reply,id=%d]", msgID)
	return r
}

func (r *Request) SendGroupMsg() ISendGroupMsg {
	cmd := SendGroupMsg
	r.Action = string(cmd)
	return r
}

func (r *Request) ToGroupID(group int64) IMsg {
	r.Params.GroupID = group
	return r
}

func (r *Request) SendPrivateMsg() ISendPrivateMsg {
	cmd := SendPrivateMsg
	r.Action = string(cmd)
	return r
}

func (r *Request) ToPrivateID(user int64) IMsg {
	r.Params.UserID = user
	return r
}

func (r *Request) TextMsg(text string) IMsg {
	r.Params.Message += text
	return r
}

func (r *Request) JsonMsg(json string) IMsg {
	r.Params.Message += fmt.Sprintf("[CQ:json,data=%v]", json)
	return r
}

func (r *Request) Face(ID int) IMsg {
	r.Params.Message += fmt.Sprintf("[CQ:face,id=%d]", ID)
	return r
}

func (r *Request) ImgMsg(img string) IMsg {
	r.Params.Message += fmt.Sprintf("[CQ:image,file=%v]", img)
	return r
}
func (r *Request) ImgBase64Msg(imgBase64 string) IMsg {
	r.Params.Message += fmt.Sprintf("[CQ:image,file=base64://%v]", imgBase64)
	return r
}
