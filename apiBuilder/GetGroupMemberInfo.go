package apiBuilder

type IGroupMemberInfo interface {
	ToGroupIDAndUserID(groupID int64, userID int64) IGroupMemberInfo
	DoApi
}

func (r *Request) GetGroupMemberInfo() IGroupMemberInfo {
	return r
}

func (r *Request) ToGroupIDAndUserID(groupID int64, userID int64) IGroupMemberInfo {
	r.Action = string(GetGroupMemberInfo)
	r.Params.GroupID = groupID
	r.Params.UserID = userID
	return r
}
