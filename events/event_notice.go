package events

type NoticeType string

// 通知类型
const (
	NoticeGroupUpload   NoticeType = "group_upload"
	NoticeGroupAdmin    NoticeType = "group_admin"
	NoticeGroupDecrease NoticeType = "group_decrease"
	NoticeGroupIncrease NoticeType = "group_increase"
	NoticeGroupBan      NoticeType = "group_ban"
	NoticeFriendAdd     NoticeType = "friend_add"
	NoticeGroupReCall   NoticeType = "group_recall"
	NoticeFriendReCall  NoticeType = "friend_recall"
	NoticeNotify        NoticeType = "notify" // 戳一戳 群红包运气王 群成员荣誉变更
)

/*
*
接收EVENT
子类型为空, 原始通知类型
*/
const (
	EventGroupReCall  EventName = "group_recall"
	EventFriendReCall EventName = "friend_recall"

	GroupUpload EventName = "group_upload" // 通知事件
	GroupNotify EventName = "notify"

	FriendAdd EventName = "friend_add"

	EventSetAdmin   EventName = "set" // 事件子类型
	EventUnSetAdmin EventName = "unset"

	EventLeave  EventName = "leave" // 事件子类型，分别表示主动退群、成员被踢、登录号被踢
	EventKick   EventName = "kick"
	EventKickMe EventName = "kick_me"

	EventInvite  EventName = "invite" // 事件子类型，分别表示管理员已同意入群、管理员邀请入群
	EventApprove EventName = "approve"

	EventBan     EventName = "ban" // 事件子类型，分别表示禁言、解除禁言
	EventLiftBan EventName = "lift_ban"

	EventPoke      EventName = "poke"
	EventLuckyKing EventName = "lucky_king"
	EventHonor     EventName = "honor"
)

type ICommonNotice interface {
	GetNoticeType() NoticeType
	GetNoticeSubType() string
	GetNoticeUserID() int64
	GetEventNoticeStruct() EventNoticeStruct
}

type IGroupNotice interface {
	GetNoticeGroupID() int64
}

// IGroupUpload 群上传文件
type IGroupUpload interface {
	IGroupNotice
	ICommonNotice
	GetFile() File
}

// IGroupAdmin 群管理员变动
type IGroupAdmin interface {
	ICommonNotice
	IGroupNotice
}

// IGroupDecrease 群成员减少
type IGroupDecrease interface {
	ICommonNotice
	IGroupNotice
	GetOperatorID() int64
}

// IGroupIncrease 群成员增加
type IGroupIncrease interface {
	ICommonNotice
	IGroupNotice
	GetOperatorID() int64
}

// IGroupBan 群禁言
type IGroupBan interface {
	ICommonNotice
	IGroupNotice
	GetOperatorID() int64
	GetDuration() int64
}

// IFriendAdd 好友添加
type IFriendAdd interface {
	ICommonNotice
}

// IGroupReCall 群消息撤回
type IGroupReCall interface {
	ICommonNotice
	IGroupNotice
	GetOperatorID() int64
	GetNoticeMessageID() int64
}

// IFriendReCall 好友消息撤回
type IFriendReCall interface {
	ICommonNotice
	GetNoticeMessageID() int64
}

// IPoke 戳一戳
type IPoke interface {
	ICommonNotice
	IGroupNotice
	GetTargetID() int64
}

// ILuckyKing 群红包运气王
type ILuckyKing interface {
	ICommonNotice
	IGroupNotice
	GetTargetID() int64
}

// IHonor 群成员荣誉变更
type IHonor interface {
	ICommonNotice
	IGroupNotice
	GetHonorType() string
}

type EventNoticeStruct struct {
	Time   int64 `json:"time"`    // 事件发生的时间戳
	SelfID int64 `json:"self_id"` // 收到事件的机器人 QQ 号

	PostType   PostType   `json:"post_type"`   // 上报类型
	NoticeType NoticeType `json:"notice_type"` // 通知类型
	SubType    string     `json:"sub_type"`    // 事件子类型
	HonorType  string     `json:"honor_type"`  // 荣誉类型，分别表示龙王、群聊之火、快乐源泉

	GroupID    int64 `json:"group_id"` // 群号
	MessageID  int64 `json:"message_id"`
	OperatorID int64 `json:"operator_id"` // 操作者 QQ号（如果是主动退群，则和 user_id 相同）
	UserID     int64 `json:"user_id"`     // 发送者QQ号
	File       File  `json:"file"`        // 文件信息
	Duration   int64 `json:"duration"`    // 时长
	TargetID   int64 `json:"target_id"`   // 被戳者 QQ号

}

type File struct {
	ID    string `json:"id,omitempty"`   // 文件 ID
	Name  string `json:"name,omitempty"` // 文件名
	Size  int64  `json:"size,omitempty"` // 文件大小 (字节数)
	Busid int64  `json:"busid,omitempty"`
}

func (e *Event) GetNoticeType() NoticeType {
	return e.EventNoticeStruct.NoticeType
}

func (e *Event) GetNoticeSubType() string {
	return e.EventNoticeStruct.SubType
}

func (e *Event) GetEventNoticeStruct() EventNoticeStruct {
	return e.EventNoticeStruct
}

func (e *Event) GetNoticeMessageID() int64 {
	return e.EventNoticeStruct.MessageID
}

func (e *Event) GetNoticeUserID() int64 {
	return e.EventNoticeStruct.UserID
}

func (e *Event) GetNoticeGroupID() int64 {
	return e.EventNoticeStruct.GroupID
}

func (e *Event) GetFile() File {
	return e.EventNoticeStruct.File
}

func (e *Event) GetOperatorID() int64 {
	return e.EventNoticeStruct.OperatorID
}

func (e *Event) GetDuration() int64 {
	return e.EventNoticeStruct.Duration
}

func (e *Event) GetTargetID() int64 {
	return e.EventNoticeStruct.TargetID
}

func (e *Event) GetHonorType() string {
	return e.EventNoticeStruct.HonorType
}
