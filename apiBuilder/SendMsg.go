package apiBuilder

type ISendGroupMsg interface {
	IMsg
}

type IMsg interface {
	ToGroupID(group int64) IMsg
	TextMsg(text string) IMsg
	DoApi
}

func (r *Request) SendGroupMsg() ISendGroupMsg {
	cmd := "send_group_msg"
	r.Action = cmd
	return r
}

func (r *Request) ToGroupID(group int64) IMsg {
	r.Params.GroupID = group
	return r
}

func (r *Request) TextMsg(text string) IMsg {
	r.Params.Message = text
	return r
}
