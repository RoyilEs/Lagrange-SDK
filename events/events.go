package events

import (
	"context"
	"encoding/json"
)

//go:generate easyjson events.go

type EventName string

const (
	EventGroupMsg    EventName = "group"
	EventPrivateMsg  EventName = "private"
	EventSetAdmin    EventName = "set"
	EventUnSetAdmin  EventName = "unset"
	EventInvite      EventName = "invite"
	EventApprove     EventName = "approve"
	EventKick        EventName = "kick"
	EventKickMe      EventName = "kick_me"
	EventLeave       EventName = "leave"
	EventGroupReCall EventName = "group_recall"
)

type PostType string

const (
	MESSAGE   PostType = "message"
	NOTICE    PostType = "notice"
	REQUEST   PostType = "request"
	METAEVENT PostType = "meta_event"
)

type EventCallbackFunc func(ctx context.Context, event IEvent)

// IEvent TODO 整合onebot四种推送事件 每个Api所推送回的信息体
type IEvent interface {
	IEventMessage
	IEventNotice
	IEVentStatus
}

type IEventMessage interface {
	ICommonMsg
	ParseGroupMsg() IGroupMsg
	ParsePrivateMsg() IPrivateMsg
}

type IEventNotice interface {
	ParseSet() ISet
	ParseUnSet() IUnSet
	ParseInvite() IInvite
	ParseKick() IKick
	ParseLeave() IKick
	ParseGroupReCall() IGroupReCall
}

type IEVentStatus interface {
	ParseGroupMemberInfo() IGroupMemberInfo
	ParseGroupMsgInfo() IGroupMsgID
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
		ok = event.EventNoticeStruct.GetSubType()
		if ok == "" {
			ok = event.EventNoticeStruct.GetNoticeType()
		}
	case string(REQUEST):
		ok = event.EventNoticeStruct.GetSubType()
	case string(METAEVENT):
		ok = event.EventNoticeStruct.GetSubType()
	}
	if event.EventStatus.Status == "ok" {
		err = json.Unmarshal(data, &event.EventStatus)
	}
	return event, data, ok, nil
}

type Event struct {
	rawEvent           []byte
	EventMessageStruct EventMessageStruct
	EventNoticeStruct  EventNoticeStruct
	EventStatus        EventStatus
	EventStruct
}

type EventStruct struct {
	Time      int64  `json:"time"`
	SelfID    int64  `json:"self_id"`
	PostType  string `json:"post_type"`
	CurrentQQ int64  `json:"current_qq"`
}

/**
IEventMessage
*/

func (e *Event) ParseGroupMsg() IGroupMsg {
	return e
}

func (e *Event) ParsePrivateMsg() IPrivateMsg {
	return e
}

/**
IEventNotice
*/

func (e *Event) ParseSet() ISet {
	return &e.EventNoticeStruct
}

func (e *Event) ParseUnSet() IUnSet {
	return &e.EventNoticeStruct
}

func (e *Event) ParseInvite() IInvite {
	return &e.EventNoticeStruct
}

func (e *Event) ParseKick() IKick {
	return &e.EventNoticeStruct
}

func (e *Event) ParseLeave() IKick {
	return &e.EventNoticeStruct
}

func (e *Event) ParseGroupReCall() IGroupReCall {
	return &e.EventNoticeStruct
}

/**
IEventStatus
*/

func (e *Event) ParseGroupMemberInfo() IGroupMemberInfo {
	return e
}

func (e *Event) GetForwardID() string {
	return e.EventStatus.GetForwardID()
}

func (e *Event) GetJoinTime() int64 {
	return e.EventStatus.GetJoinTime()
}

func (e *Event) GetLastSentTime() int64 {
	return e.EventStatus.GetLastSentTime()
}

func (e *Event) GetUnFriendly() bool {
	return e.EventStatus.GetUnFriendly()
}

func (e *Event) ParseGroupMsgInfo() IGroupMsgID {
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
func (e *EventStruct) GetCurrentQQ() int64 {
	return e.CurrentQQ
}
