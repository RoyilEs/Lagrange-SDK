package apiBuilder

type IGroupMemberInfo interface {
	ToGroupIDAndUserID(groupID int64, userID int64) IGroupMemberInfo
	DoApi
}

func (b *Builder) GetGroupMemberInfo() IGroupMemberInfo {
	return b
}

func (b *Builder) ToGroupIDAndUserID(groupID int64, userID int64) IGroupMemberInfo {
	b.action = GetGroupMemberInfo
	b.Params.GroupID = groupID
	b.Params.UserID = userID
	return b
}
