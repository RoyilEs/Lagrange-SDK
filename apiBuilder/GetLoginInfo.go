package apiBuilder

import "context"

type LoginInfoStruct struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nickname"`
}

func (b *Builder) GetLoginInfo(ctx context.Context) LoginInfoStruct {
	b.action = GetLoginInfo
	resp, err := b.DoResponse(ctx)
	if err != nil {
		return LoginInfoStruct{}
	}
	return LoginInfoStruct{
		UserID:   resp.Data.UserID,
		NickName: resp.Data.NickName,
	}
}
