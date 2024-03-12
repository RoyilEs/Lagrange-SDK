package apiBuilder

import (
	"Lagrange-SDK/message"
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
	LongMsg(ID string) IMsg
	Face(ID int) IMsg
	DoApi
}

type MessageStruct struct {
	Type string      `json:"type"`
	Data DataMessage `json:"data"`
}

type DataMessage struct {
	Text    string           `json:"text,omitempty"`
	File    string           `json:"file,omitempty"`
	Data    string           `json:"data,omitempty"`
	Url     string           `json:"url,omitempty"`
	QQ      string           `json:"qq,omitempty"`
	ID      string           `json:"id,omitempty"`
	Name    string           `json:"name,omitempty"`
	Uin     string           `json:"uin,omitempty"`
	Content []*MarkDownBuild `json:"content,omitempty"`
}

type MarkDownBuild struct {
	Type string `json:"type,omitempty"`
	Data struct {
		Content string `json:"content,omitempty"`
	} `json:"data,omitempty"`
}

func (b *Builder) SendReply(msgID int64) ISendReply {
	i := append(b.Params.Message, &MessageStruct{
		Type: message.REPLY,
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
	i := append(b.Params.Message, &MessageStruct{
		Type: message.TEXT,
		Data: DataMessage{
			Text: text,
		},
	})
	b.Params.Message = i
	return b
}

func (b *Builder) JsonMsg(json string) IMsg {
	i := append(b.Params.Message, &MessageStruct{
		Type: message.JSON,
		Data: DataMessage{
			Data: json,
		},
	})
	b.Params.Message = i
	return b

}

func (b *Builder) Face(ID int) IMsg {
	i := append(b.Params.Message, &MessageStruct{
		Type: message.FACE,
		Data: DataMessage{
			ID: strconv.Itoa(ID),
		},
	})
	b.Params.Message = i
	return b
}

func (b *Builder) ImgMsg(img string) IMsg {
	i := append(b.Params.Message, &MessageStruct{
		Type: message.IMAGE,
		Data: DataMessage{
			File: img,
		},
	})
	b.Params.Message = i
	return b
}
func (b *Builder) ImgBase64Msg(imgBase64 string) IMsg {
	i := append(b.Params.Message, &MessageStruct{
		Type: message.IMAGE,
		Data: DataMessage{
			File: "base64://" + imgBase64,
		},
	})
	b.Params.Message = i
	return b
}

func (b *Builder) LongMsg(ID string) IMsg {
	i := append(b.Params.Message, &MessageStruct{
		Type: message.LONGMSG,
		Data: DataMessage{
			ID: ID,
		},
	})
	b.Params.Message = i
	return b
}
