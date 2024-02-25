package events

type EventNoticeStruct struct {
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	NoticeType string `json:"notice_type"`
}
