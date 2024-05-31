package events

type ISet interface {
	ICommonNotice
}

type IUnSet interface {
	ICommonNotice
}

type IInvite interface {
	ICommonNotice
}

type IKick interface {
	ICommonNotice
}

type IGroupReCall interface {
	ICommonNotice
}

type ICommonNotice interface {
	GetSubType() string
	GetGroupID() int64
	GetOperatorID() int64
	GetMessageID() int64
	GetUserID() int64
	GetNoticeType() string
}

type EventNoticeStruct struct {
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	MessageID  int64  `json:"message_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
	NoticeType string `json:"notice_type"`
}

func (e *EventNoticeStruct) GetSubType() string {
	return e.SubType
}

func (e *EventNoticeStruct) GetGroupID() int64 {
	return e.GroupID
}

func (e *EventNoticeStruct) GetOperatorID() int64 {
	return e.OperatorID
}

func (e *EventNoticeStruct) GetMessageID() int64 {
	return e.MessageID
}

func (e *EventNoticeStruct) GetUserID() int64 {
	return e.UserID
}

func (e *EventNoticeStruct) GetNoticeType() string {
	return e.NoticeType
}
