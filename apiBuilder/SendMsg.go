package apiBuilder

import (
	"strconv"
)

type ISendReply interface {
	SendGroupMsg(groupID int64) ISendGroupMsg
	SendPrivateMsg(userID int64) ISendPrivateMsg
}

type ISendGroupMsg interface {
	IMsg
	ISendReply
}

type ISendPrivateMsg interface {
	IMsg
	ISendReply
}

type IMsg interface {
	TextMsg(text string) IMsg
	JsonMsg(json string) IMsg
	ImgMsg(img string) IMsg
	ImgBase64Msg(imgBase64 string) IMsg
	Face(ID int) IMsg
	DoApi
}

type MessageStruct struct {
	Type string      `json:"type"`
	Data DataMessage `json:"data"`
}

type DataMessage struct {
	Text string `json:"text,omitempty"`
	File string `json:"file,omitempty"`
	Data string `json:"data,omitempty"`
	Url  string `json:"url,omitempty"`
	QQ   string `json:"qq,omitempty"`
	ID   string `json:"id,omitempty"`
}

func (b *Builder) SendReply(msgID int64) ISendReply {
	i := append(b.Params.Message, MessageStruct{
		Type: "reply",
		Data: DataMessage{
			ID: strconv.FormatInt(msgID, 10),
		},
	})
	b.Params.Message = i
	return b
}

func (b *Builder) SendGroupMsg(GroupID int64) ISendGroupMsg {
	cmd := SendGroupMsg
	b.action = cmd
	b.Params.GroupID = GroupID
	return b
}

func (b *Builder) SendPrivateMsg(userID int64) ISendPrivateMsg {
	cmd := SendPrivateMsg
	b.action = cmd
	b.Params.UserID = userID
	return b
}

func (b *Builder) TextMsg(text string) IMsg {
	i := append(b.Params.Message, MessageStruct{
		Type: "text",
		Data: DataMessage{
			Text: text,
		},
	})
	b.Params.Message = i
	return b
}

func (b *Builder) JsonMsg(json string) IMsg {
	i := append(b.Params.Message, MessageStruct{
		Type: "json",
		Data: DataMessage{
			Data: json,
		},
	})
	b.Params.Message = i
	return b

}

func (b *Builder) Face(ID int) IMsg {
	i := append(b.Params.Message, MessageStruct{
		Type: "face",
		Data: DataMessage{
			ID: strconv.Itoa(ID),
		},
	})
	b.Params.Message = i
	return b
}

func (b *Builder) ImgMsg(img string) IMsg {
	i := append(b.Params.Message, MessageStruct{
		Type: "image",
		Data: DataMessage{
			File: img,
		},
	})
	b.Params.Message = i
	return b
}
func (b *Builder) ImgBase64Msg(imgBase64 string) IMsg {
	i := append(b.Params.Message, MessageStruct{
		Type: "image",
		Data: DataMessage{
			File: "base64://" + imgBase64,
		},
	})
	b.Params.Message = i
	return b
}
