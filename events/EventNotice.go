package events

type ISet interface {
	ICommonNotice
}

type IUnSet interface {
	ICommonNotice
}

type ICommonNotice interface {
	GetSubType() string
	GetGroupID() int64
	GetUserID() int64
	GetNoticeType() string
}

type EventNoticeStruct struct {
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	NoticeType string `json:"notice_type"`
}

func (e *EventNoticeStruct) GetSubType() string {
	return e.SubType
}

func (e *EventNoticeStruct) GetGroupID() int64 {
	return e.GroupID
}

func (e *EventNoticeStruct) GetUserID() int64 {
	return e.UserID
}

func (e *EventNoticeStruct) GetNoticeType() string {
	return e.NoticeType
}
