package apiBuilder

type ISetGroupKick interface {
	ToGroupIDAndKickUserID(groupID int64, userID int64) ISetGroupBan
	DoApi
}

func (r *Request) SetGroupKick() ISetGroupKick {
	return r
}

func (r *Request) ToGroupIDAndKickUserID(groupID int64, userID int64) ISetGroupBan {
	r.Action = string(SetGroupKick)
	r.Params.GroupID = groupID
	r.Params.UserID = userID
	return r
}
