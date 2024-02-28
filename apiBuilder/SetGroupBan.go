package apiBuilder

type ISetGroupBan interface {
	ToGroupIDAndMuteUserID(groupID int64, userID int64) ISetGroupBan
	Duration(duration int) ISetGroupBan
	DoApi
}

func (r *Request) SetGroupBan() ISetGroupBan {
	return r
}

func (r *Request) ToGroupIDAndMuteUserID(groupID int64, userID int64) ISetGroupBan {
	r.Action = string(SetGroupBan)
	r.Params.GroupID = groupID
	r.Params.UserID = userID
	return r
}

func (r *Request) Duration(duration int) ISetGroupBan {
	r.Params.Duration = duration
	return r
}
