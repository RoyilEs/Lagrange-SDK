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
	EventSetAdmin   EventName = "set"
	EventUnSetAdmin EventName = "unset"
)

type PostType string

const (
	MESSAGE   PostType = "message"
	NOTICE    PostType = "notice"
	REQUEST   PostType = "request"
	METAEVENT PostType = "meta_event"
)

type EventCallbackFunc func(client *websocket.Conn, event IEvent)

// IEvent TODO 整合onebot四种推送事件
type IEvent interface {
	IEventMessage
	IEventNotice
}

type IEventMessage interface {
	ICommonMsg
	ParseGroupMsg() IGroupMsg
	ParsePrivateMsg() IPrivateMsg
}

type IEventNotice interface {
}

func New(data []byte) (*Event, []byte, string, error) {
	event := &Event{}
	err := json.Unmarshal(data, event)
	if err != nil {
		return nil, nil, "", nil
	}
	event.rawEvent = data
	var ok string
	switch event.GetPostType() {
	case string(MESSAGE):
		err = json.Unmarshal(data, &event.EventMessageStruct)
		ok = event.GetMessageType()
	case string(NOTICE):
		err = json.Unmarshal(data, &event.EventNoticeStruct)
		ok = event.GetNoticeSubType()
	case string(REQUEST):
		ok = event.GetNoticeSubType()
	case string(METAEVENT):
		ok = event.GetNoticeSubType()
	}
	return event, data, ok, nil
}

type Event struct {
	rawEvent           []byte
	EventMessageStruct EventMessageStruct
	EventNoticeStruct  EventNoticeStruct
	EventStruct
}

type EventStruct struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}

func (e *Event) ParseGroupMsg() IGroupMsg {
	return e
}

func (e *Event) ParsePrivateMsg() IPrivateMsg {
	return e
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
