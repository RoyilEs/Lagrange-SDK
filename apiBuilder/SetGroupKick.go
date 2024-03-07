package apiBuilder

type ISetGroupKick interface {
	ToGroupIDAndKickUserID(groupID int64, userID int64) ISetGroupBan
	DoApi
}

func (b *Builder) SetGroupKick() ISetGroupKick {
	return b
}

func (b *Builder) ToGroupIDAndKickUserID(groupID int64, userID int64) ISetGroupBan {
	b.action = SetGroupKick
	b.Params.GroupID = groupID
	b.Params.UserID = userID
	return b
}
