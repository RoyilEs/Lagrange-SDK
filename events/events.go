package events

import (
	"context"
	"encoding/json"
)

type EventName string

type PostType string

const (
	MESSAGE   PostType = "message"
	NOTICE    PostType = "notice"
	REQUEST   PostType = "request"
	METAEVENT PostType = "meta_event"
)

type EventCallbackFunc func(ctx context.Context, event IEvent)

type IEvent interface {
	Message() IEventMessage
	Notice() IEventNotice
	Request() IEventRequest
	MetaEvent() IMetaEvent
}

func (e *Event) Message() IEventMessage {
	return e
}
func (e *Event) Notice() IEventNotice {
	return e
}
func (e *Event) Request() IEventRequest {
	return e
}
func (e *Event) MetaEvent() IMetaEvent {
	return e
}

// IEventMessage 信息事件方法
type IEventMessage interface {
	ICommonMsg
	ParseGroupMsg() IGroupMessage
	ParsePrivateMsg() IPrivateMessage
}

// IEventNotice 通知事件方法
type IEventNotice interface {
	ICommonNotice
	ParseGroupUpload() IGroupUpload

	ParseGroupSetAdmin() IGroupAdmin
	ParseGroupUnSetAdmin() IGroupAdmin

	ParseGroupLeave() IGroupDecrease
	ParseGroupKick() IGroupDecrease
	ParseGroupKickMe() IGroupDecrease

	ParseGroupApprove() IGroupIncrease
	ParseGroupInvite() IGroupIncrease

	ParseGroupBan() IGroupBan
	ParseGroupLiftBan() IGroupBan

	ParseFriendAdd() IFriendAdd

	ParseGroupReCall() IGroupReCall
	ParseFriendReCall() IFriendReCall

	ParsePoke() IPoke
	ParseLuckyKing() ILuckyKing
	ParseHonor() IHonor
}

// IEventRequest 请求事件方法
type IEventRequest interface {
	ICommonRequest
	ParseFriend() IRequestFriend
	ParseGroup() IRequestGroup
}

type IMetaEvent interface {
	ICommonMeta
	Lifecycle() ILifeCycle
	Heartbeat() IHeartbeat
}

type EventStruct struct {
	Time      int64  `json:"time"`
	SelfID    int64  `json:"self_id"`
	PostType  string `json:"post_type"`
	CurrentQQ int64  `json:"current_qq"`
}

type Event struct {
	rawEvent           []byte
	EventMessageStruct EventMessageStruct
	EventNoticeStruct  EventNoticeStruct
	EventRequestStruct EventRequestStruct
	MetaEventStruct    MetaEventStruct
	EventStatus        EventStatus
	EventStruct
}

type EventStatus struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
	Data    Data   `json:"data,omitempty"`
	Echo    any    `json:"echo"`
}

type Data struct {
	GroupID      int64  `json:"group_id,omitempty"`
	UserID       int64  `json:"user_id,omitempty"`
	MessageID    int64  `json:"message_id,omitempty"`
	ForwardID    string `json:"forward_id,omitempty"`
	NickName     string `json:"nickname,omitempty"`
	Card         any    `json:"card,omitempty"`
	Sex          string `json:"sex,omitempty"`
	Age          int    `json:"age,omitempty"`
	Area         string `json:"area,omitempty"`
	JoinTime     int64  `json:"join_time,omitempty"`
	LastSentTime int64  `json:"last_sent_time,omitempty"`
	Level        string `json:"level,omitempty"`
	Role         string `json:"role,omitempty"`
	UnFriendly   bool   `json:"unfriendly,omitempty"`
	Title        string `json:"title,omitempty"`
}

func New(data []byte) (*Event, []byte, string, error) {
	event := &Event{}
	err := json.Unmarshal(data, event)
	if err != nil {
		return nil, nil, "", err
	}
	event.rawEvent = data
	var ok string
	switch event.PostType {
	case string(MESSAGE):
		err = json.Unmarshal(data, &event.EventMessageStruct)
		ok = event.EventMessageStruct.MessageType
	case string(NOTICE):
		err = json.Unmarshal(data, &event.EventNoticeStruct)
		if ok = event.EventNoticeStruct.SubType; ok == "" {
			ok = string(event.EventNoticeStruct.NoticeType)
		}
	case string(REQUEST):
		err = json.Unmarshal(data, &event.EventRequestStruct)
		if ok = event.EventRequestStruct.SubType; ok == "" {
			ok = string(event.EventRequestStruct.RequestType)
		}
	case string(METAEVENT):
		err = json.Unmarshal(data, &event.MetaEventStruct)
		ok = string(event.MetaEventStruct.MetaEventType)
	}
	if event.EventStatus.Status == "ok" {
		err = json.Unmarshal(data, &event.EventStatus)
	}
	return event, data, ok, err
}

func (e *Event) ParseGroupMsg() IGroupMessage {
	return e
}

func (e *Event) ParsePrivateMsg() IPrivateMessage {
	return e
}

func (e *Event) ParseGroupUpload() IGroupUpload {
	return e
}

func (e *Event) ParseGroupSetAdmin() IGroupAdmin {
	return e
}

func (e *Event) ParseGroupUnSetAdmin() IGroupAdmin {
	return e
}

func (e *Event) ParseGroupLeave() IGroupDecrease {
	return e
}

func (e *Event) ParseGroupKick() IGroupDecrease {
	return e
}

func (e *Event) ParseGroupKickMe() IGroupDecrease {
	return e
}

func (e *Event) ParseGroupApprove() IGroupIncrease {
	return e
}

func (e *Event) ParseGroupInvite() IGroupIncrease {
	return e
}

func (e *Event) ParseGroupBan() IGroupBan {
	return e
}

func (e *Event) ParseGroupLiftBan() IGroupBan {
	return e
}

func (e *Event) ParseFriendAdd() IFriendAdd {
	return e
}

func (e *Event) ParseGroupReCall() IGroupReCall {
	return e
}

func (e *Event) ParseFriendReCall() IFriendReCall {
	return e
}

func (e *Event) ParsePoke() IPoke {
	return e
}

func (e *Event) ParseLuckyKing() ILuckyKing {
	return e
}

func (e *Event) ParseHonor() IHonor {
	return e
}

func (e *Event) ParseFriend() IRequestFriend {
	return e
}

func (e *Event) ParseGroup() IRequestGroup {
	return e
}

func (e *Event) Lifecycle() ILifeCycle {
	return e
}

func (e *Event) Heartbeat() IHeartbeat {
	return e
}
