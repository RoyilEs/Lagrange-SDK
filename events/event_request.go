package events

type RequestType string

// 通知类型
const (
	RequestFriend RequestType = "friend"
	RequestGroup  RequestType = "group"
)

/*
*
接收EVENT
子类型为空, 原始通知类型
*/
const (
	EventFriendRequest EventName = "friend"
	EventGroupRequest  EventName = "group"

	EventADD EventName = "add" // invite已有直接调用
)

type ICommonRequest interface {
	GetRequestPostType() PostType
	GetRequestType() RequestType
	GetRequestUserID() int64
	GetComment() string
	GetFlag() string
}

type IRequestFriend interface {
	ICommonRequest
}

type IRequestGroup interface {
	ICommonRequest
	GetRequestSubType() string
	GetRequestGroupID() int64
}

type EventRequestStruct struct {
	Time   int64 `json:"time"`    // 事件发生的时间戳
	SelfID int64 `json:"self_id"` // 收到事件的机器人 QQ 号

	PostType    PostType    `json:"post_type"`    // 上报类型
	RequestType RequestType `json:"request_type"` // 通知类型
	SubType     string      `json:"sub_type"`     // 事件子类型

	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"` // 验证信息
	Flag    string `json:"flag"`    // 请求 flag，在调用处理请求的 API 时需要传入
}

func (e *Event) GetRequestPostType() PostType {
	return e.EventRequestStruct.PostType
}

func (e *Event) GetRequestType() RequestType {
	return e.EventRequestStruct.RequestType
}

func (e *Event) GetComment() string {
	return e.EventRequestStruct.Comment
}

func (e *Event) GetFlag() string {
	return e.EventRequestStruct.Flag
}

func (e *Event) GetRequestSubType() string {
	return e.EventRequestStruct.SubType
}

func (e *Event) GetRequestGroupID() int64 {
	return e.EventRequestStruct.GroupID
}

func (e *Event) GetRequestUserID() int64 {
	return e.EventRequestStruct.UserID
}
