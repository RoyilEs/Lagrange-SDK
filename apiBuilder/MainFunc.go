package apiBuilder

type ApiName string

const (
	GetGroupMemberInfo ApiName = "get_group_member_info"
	SetGroupBan        ApiName = "set_group_ban"
	SendGroupMsg       ApiName = "send_group_msg"
	SendPrivateMsg     ApiName = "send_private_msg"
)

type IMainFuncApi interface {
	SendGroupMsg() ISendGroupMsg
	SendPrivateMsg() ISendPrivateMsg
	GetGroupMemberInfo() IGroupMemberInfo
	SetGroupBan() ISetGroupBan
}
