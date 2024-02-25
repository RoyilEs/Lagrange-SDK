package apiBuilder

import (
	"fmt"
)

type ISendGroupMsg interface {
	IMsg
	ToGroupID(group int64) IMsg
}

type ISendPrivateMsg interface {
	IMsg
	ToPrivateID(user int64) IMsg
}

type IMsg interface {
	TextMsg(text string) IMsg
	ImgMsg(img string) IMsg
	Face(ID int) IMsg
	DoApi
}

func (r *Request) SendGroupMsg() ISendGroupMsg {
	cmd := "send_group_msg"
	r.Action = cmd
	return r
}

func (r *Request) ToGroupID(group int64) IMsg {
	r.Params.GroupID = group
	return r
}

func (r *Request) SendPrivateMsg() ISendPrivateMsg {
	cmd := "send_private_msg"
	r.Action = cmd
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

func (r *Request) Face(ID int) IMsg {
	r.Params.Message += fmt.Sprintf("[CQ:face,id=%d]", ID)
	return r
}

func (r *Request) ImgMsg(img string) IMsg {
	r.Params.Message += fmt.Sprintf("[CQ:image,file=%v]", img)
	return r
}
