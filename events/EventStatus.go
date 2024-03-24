package events

type IGroupMsgID interface {
	GetMessageID() int64
}

type IGroupMemberInfo interface {
	GetGroupID() int64
	GetUserID() int64
	GetMessageID() int64
	GetForwardID() string
	GetNickName() string
	GetCard() string
	GetSex() string
	GetAge() int
	GetArea() string
	GetJoinTime() int64
	GetLastSentTime() int64
	GetLevel() string
	GetRole() string
	GetUnFriendly() bool
	GetTitle() string
}

func (e *EventStatus) GetGroupID() int64 {
	return e.Data.GroupID
}

func (e *EventStatus) GetUserID() int64 {
	return e.Data.UserID
}

func (e *EventStatus) GetMessageID() int64 {
	return e.Data.MessageID
}

func (e *EventStatus) GetForwardID() string {
	return e.Data.ForwardID
}

func (e *EventStatus) GetNickName() string {
	return e.Data.NickName
}

func (e *EventStatus) GetCard() any {
	return e.Data.Card
}

func (e *EventStatus) GetSex() string {
	return e.Data.Sex
}

func (e *EventStatus) GetAge() int {
	return e.Data.Age
}

func (e *EventStatus) GetArea() string {
	return e.Data.Area
}

func (e *EventStatus) GetJoinTime() int64 {
	return e.Data.JoinTime
}

func (e *EventStatus) GetLastSentTime() int64 {
	return e.Data.LastSentTime
}

func (e *EventStatus) GetLevel() string {
	return e.Data.Level
}

func (e *EventStatus) GetRole() string {
	return e.Data.Role
}

func (e *EventStatus) GetUnFriendly() bool {
	return e.Data.UnFriendly
}

func (e *EventStatus) GetTitle() string {
	return e.Data.Title
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
