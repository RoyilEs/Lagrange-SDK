package apiBuilder

import "context"

type ApiName string

const (
	GetGroupMemberInfo  ApiName = "get_group_member_info"
	GetLoginInfo        ApiName = "get_login_info"
	SetGroupKick        ApiName = "set_group_kick"
	SetGroupBan         ApiName = "set_group_ban"
	SendGroupMsg        ApiName = "send_group_msg"
	SendPrivateMsg      ApiName = "send_private_msg"
	SendGroupForwardMsg ApiName = "send_group_forward_msg"
)

type IMainFuncApi interface {
	SendReply(msgID int64) ISendReply
	SendGroupMsg(groupID int64) ISendGroupMsg
	SendPrivateMsg(userID int64) ISendPrivateMsg
	GetGroupMemberInfo() IGroupMemberInfo
	GetLoginInfo(ctx context.Context) LoginInfoStruct
	SetGroupBan() ISetGroupBan
	SetGroupKick() ISetGroupKick
	MarkDownBuild(name string, uin int64) IMarkDownBuild
}
