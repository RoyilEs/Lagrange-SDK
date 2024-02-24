package events

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

//go:generate easyjson events.go

type EventName string

const (
	EventGroupMsg   EventName = "group"
	EventPrivateMsg EventName = "private"
)

type EventCallbackFunc func(client *websocket.Conn, event IEvent)

type IEvent interface {
	ICommonMsg
	ParseGroupMsg() IGroupMsg
}

type IPrivateMsg interface {
	ICommonMsg
	IPrivateSender
}

type IGroupMsg interface {
	ICommonMsg
	IGroupSender
	ParseTextMsg() IMessage
	GetGroupID() int64
}

type IGroupSender interface {
	IPrivateSender
	GetCard() any
	GetAge() int
	GetArea() string
	GetLevel() string
	GetRole() string
	GetTitle() string
}

type IPrivateSender interface {
	GetUserID() int64
	GetNickName() string
	GetSex() string
}
type IMessage interface {
	GetType() []string
	GetText() []string
	GetFile() []string
	GetUrl() []string
	GetQQ() []string
	GetID() []string
}

type ICommonMsg interface {
	GetMessageType() EventName
	GetSubType() string
	GetMessageID() int64
	GetUserID() int64
	GetTime() int64
	GetSelfID() int64
	GetPostType() string
}

func New(data []byte) (*EventStruct, error) {
	event := &EventStruct{}
	err := json.Unmarshal(data, event)
	if err != nil {
		return nil, err
	}
	event.rawEvent = data
	return event, nil
}

type EventStruct struct {
	rawEvent    []byte
	MessageType string `json:"message_type,omitempty"`
	SubType     string `json:"sub_type,omitempty"`
	MessageID   int64  `json:"message_id,omitempty"`
	GroupID     int64  `json:"group_id"`
	UserID      int64  `json:"user_id,omitempty"`
	Anonymous   any    `json:"anonymous"`
	Message     *[]struct {
		Type string `json:"type"`
		Data struct {
			Text string `json:"text"`
			File string `json:"file"`
			Url  string `json:"url"`
			QQ   string `json:"qq"`
			ID   string `json:"id"`
		}
	} `json:"message,omitempty"`
	RawMessage string `json:"raw_message,omitempty"`
	Font       int    `json:"font,omitempty"`
	Sender     *struct {
		UserID   int64  `json:"user_id"`
		NickName string `json:"nickname"`
		Card     any    `json:"card"`
		Sex      string `json:"sex"`
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Level    string `json:"level"`
		Role     string `json:"role"`
		Title    string `json:"title"`
	} `json:"sender,omitempty"`
	Time      int64  `json:"time"`
	SelfID    int64  `json:"self_id"`
	PostType  string `json:"post_type"`
	CurrentQQ int64  `json:"CurrentQQ"`
}

/**
ICommonMsg 部分
*/

func (e *EventStruct) GetMessageType() EventName {
	return EventName(e.MessageType)
}

func (e *EventStruct) GetSubType() string {
	return e.SubType
}

func (e *EventStruct) GetMessageID() int64 {
	return e.MessageID
}

func (e *EventStruct) GetUserID() int64 {
	return e.UserID
}

func (e *EventStruct) GetTime() int64 {
	return e.Time
}

func (e *EventStruct) GetSelfID() int64 {
	return e.SelfID
}
func (e *EventStruct) GetPostType() string {
	return e.PostType
}

/**
IEvent 部分
*/

func (e *EventStruct) ParseGroupMsg() IGroupMsg {
	return e
}

/**
IGroupMsg 部分
*/

func (e *EventStruct) ParseTextMsg() IMessage {
	return e
}

func (e *EventStruct) GetGroupID() int64 {
	return e.GroupID
}

/**
IMessage 部分
*/

func (e *EventStruct) GetType() []string {
	var msgType []string
	for _, v := range *e.Message {
		msgType = append(msgType, v.Type)
	}
	return msgType
}

func (e *EventStruct) GetText() []string {
	var msgText []string
	for _, v := range *e.Message {
		msgText = append(msgText, v.Data.Text)
	}
	return msgText
}

func (e *EventStruct) GetFile() []string {
	var msgFile []string
	for _, v := range *e.Message {
		msgFile = append(msgFile, v.Data.File)
	}
	return msgFile
}

func (e *EventStruct) GetUrl() []string {
	var msgUrl []string
	for _, v := range *e.Message {
		msgUrl = append(msgUrl, v.Data.Url)
	}
	return msgUrl
}

func (e *EventStruct) GetQQ() []string {
	var msgQQ []string
	for _, v := range *e.Message {
		msgQQ = append(msgQQ, v.Data.QQ)
	}
	return msgQQ
}

func (e *EventStruct) GetID() []string {
	var msgID []string
	for _, v := range *e.Message {
		msgID = append(msgID, v.Data.ID)
	}
	return msgID
}

/**
IPrivateSender 部分
*/

func (e *EventStruct) GetNickName() string {
	return e.Sender.NickName
}
func (e *EventStruct) GetSex() string {
	return e.Sender.Sex
}

/**
IGroupSender 部分
*/

func (e *EventStruct) GetCard() any {
	return e.Sender.Card
}

func (e *EventStruct) GetAge() int {
	return e.Sender.Age
}

func (e *EventStruct) GetArea() string {
	return e.Sender.Area
}

func (e *EventStruct) GetLevel() string {
	return e.Sender.Level
}

func (e *EventStruct) GetRole() string {
	return e.Sender.Role
}

func (e *EventStruct) GetTitle() string {
	return e.Sender.Title
}
