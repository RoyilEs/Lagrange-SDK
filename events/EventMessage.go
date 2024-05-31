package events

import (
	"strconv"
)

type IPrivateMsg interface {
	IPrivateSender
	ParseTextMsg() IMessage
	GetUserID() int64
}

type IPrivateSender interface {
	GetUserID() int64
	GetNickName() string
	GetSex() string
}

type IGroupMsg interface {
	IGroupSender
	ParseTextMsg() IMessage
	GetGroupID() int64
	IsFromBot(botQQ int64) (flag bool)
	AtBot(botQQ int64) (at bool)
}

type IGroupSender interface {
	IPrivateSender
	GetCard() string
	GetAge() int
	GetArea() string
	GetLevel() string
	GetRole() string
	GetTitle() string
}

type IMessage interface {
	GetType() []string
	GetText() []string
	GetFile() []string
	GetData() string
	GetUrl() []string
	GetQQ() []string
	GetID() []string
	GetAtQQ() []string
}

type ICommonMsg interface {
	GetMessageType() string
	GetMessageSubType() string
	GetMessageID() int64
	GetUserID() int64
	GetEventMessageStruct() EventMessageStruct
}

type EventMessageStruct struct {
	MessageType string         `json:"message_type,omitempty"`
	SubType     string         `json:"sub_type,omitempty"`
	MessageID   int64          `json:"message_id,omitempty"`
	GroupID     int64          `json:"group_id,omitempty"`
	UserID      int64          `json:"user_id,omitempty"`
	Anonymous   any            `json:"anonymous"`
	Message     *[]MessageData `json:"message,omitempty"`
	RawMessage  string         `json:"raw_message,omitempty"`
	Font        int            `json:"font,omitempty"`
	Sender      *Sender        `json:"sender,omitempty"`
	EventStruct
}

type MessageData struct {
	Type string `json:"type"`
	Data struct {
		Text string `json:"text,omitempty"`
		File string `json:"file,omitempty"`
		Data string `json:"data,omitempty"`
		Url  string `json:"url,omitempty"`
		QQ   string `json:"qq,omitempty"`
		ID   string `json:"id,omitempty"`
	}
}
type Sender struct {
	UserID   int64  `json:"user_id,omitempty"`
	NickName string `json:"nickname,omitempty"`
	Card     string `json:"card,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Age      int    `json:"age,omitempty"`
	Area     string `json:"area,omitempty"`
	Level    string `json:"level,omitempty"`
	Role     string `json:"role,omitempty"`
	Title    string `json:"title,omitempty"`
}

func (e *Event) GetMessageSubType() string {
	return e.EventMessageStruct.SubType
}

func (e *Event) GetMessageType() string {
	return e.EventMessageStruct.MessageType
}
func (e *Event) GetMessageID() int64 {
	return e.EventMessageStruct.MessageID
}

func (e *Event) GetEventMessageStruct() EventMessageStruct {
	return e.EventMessageStruct
}

func (e *Event) GetUserID() int64 {
	return e.EventMessageStruct.UserID
}

func (e *Event) GetNickName() string {
	return e.EventMessageStruct.Sender.NickName
}

func (e *Event) GetSex() string {
	return e.EventMessageStruct.Sender.Sex
}

func (e *Event) GetCard() string {
	if e.EventMessageStruct.Sender.Card == "" {
		return e.EventMessageStruct.Sender.NickName
	}
	return e.EventMessageStruct.Sender.Card
}

func (e *Event) GetAge() int {
	return e.EventMessageStruct.Sender.Age
}

func (e *Event) GetArea() string {
	return e.EventMessageStruct.Sender.Area
}

func (e *Event) GetLevel() string {
	return e.EventMessageStruct.Sender.Level
}

func (e *Event) GetRole() string {
	return e.EventMessageStruct.Sender.Role
}

func (e *Event) GetTitle() string {
	return e.EventMessageStruct.Sender.Title
}

func (e *Event) ParseTextMsg() IMessage {
	return &e.EventMessageStruct
}

func (e *Event) GetGroupID() int64 {
	return e.EventMessageStruct.GroupID
}

func (e *Event) IsFromBot(botQQ int64) (flag bool) {
	if e.EventMessageStruct.Sender.UserID == botQQ || e.EventNoticeStruct.UserID == botQQ {
		flag = true
	}
	return flag
}

func (e *Event) AtBot(botQQ int64) (at bool) {
	for _, v := range *e.EventMessageStruct.Message {
		if v.Data.QQ == strconv.FormatInt(botQQ, 10) {
			at = true
			break
		}
	}
	return at
}

func (e *EventMessageStruct) GetType() []string {
	var msgText []string
	for _, v := range *e.Message {
		msgText = append(msgText, v.Data.Text)
	}
	return msgText
}

func (e *EventMessageStruct) GetText() []string {
	var msgText []string
	for _, v := range *e.Message {
		msgText = append(msgText, v.Data.Text)
	}
	return msgText
}

func (e *EventMessageStruct) GetFile() []string {
	var msgFile []string
	for _, v := range *e.Message {
		msgFile = append(msgFile, v.Data.File)
	}
	return msgFile
}

func (e *EventMessageStruct) GetUrl() []string {
	var msgUrl []string
	for _, v := range *e.Message {
		msgUrl = append(msgUrl, v.Data.Url)
	}
	return msgUrl
}
func (e *EventMessageStruct) GetData() string {
	var data string
	for _, v := range *e.Message {
		data += v.Data.Data
	}
	return data
}

func (e *EventMessageStruct) GetQQ() []string {
	var msgQQ []string
	for _, v := range *e.Message {
		msgQQ = append(msgQQ, v.Data.QQ)
	}
	return msgQQ
}

func (e *EventMessageStruct) GetAtQQ() []string {
	var msgAtQQ []string
	for _, v := range *e.Message {
		if v.Type == "at" {
			msgAtQQ = append(msgAtQQ, v.Data.QQ)
		}
	}
	return msgAtQQ
}

func (e *EventMessageStruct) GetID() []string {
	var msgID []string
	for _, v := range *e.Message {
		msgID = append(msgID, v.Data.ID)
	}
	return msgID
}
