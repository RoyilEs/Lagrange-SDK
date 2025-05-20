package apiBuilder

type IGroupPoke interface {
	SetGroupPoke(groupID, userID int64) IGroupPoke
	DoApi
}

func (b *Builder) Poke() IGroupPoke {
	return b
}
func (b *Builder) SetGroupPoke(groupID, userID int64) IGroupPoke {
	b.action = GroupPoke
	b.Params.GroupID = groupID
	b.Params.UserID = userID
	return b
}
