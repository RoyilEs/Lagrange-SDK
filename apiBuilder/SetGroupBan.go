package apiBuilder

type ISetGroupBan interface {
	ToGroupIDAndMuteUserID(groupID int64, userID int64) ISetGroupBan
	Duration(duration int) ISetGroupBan
	DoApi
}

func (b *Builder) SetGroupBan() ISetGroupBan {
	return b
}

func (b *Builder) ToGroupIDAndMuteUserID(groupID int64, userID int64) ISetGroupBan {
	b.action = SetGroupBan
	b.Params.GroupID = groupID
	b.Params.UserID = userID
	return b
}

func (b *Builder) Duration(duration int) ISetGroupBan {
	b.Params.Duration = duration
	return b
}
