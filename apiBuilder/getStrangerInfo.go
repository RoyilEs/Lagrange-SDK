package apiBuilder

import "context"

type StrangerInfoStruct struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
}

func (b *Builder) GetStrangerInfo(ctx context.Context, userID int64) StrangerInfoStruct {
	b.action = GetStrangerInfo
	b.Params.UserID = userID
	resp, err := b.DoResponse(ctx)
	if err != nil {
		return StrangerInfoStruct{}
	}
	return StrangerInfoStruct{
		UserID:   resp.Data.UserID,
		NickName: resp.Data.NickName,
		Sex:      resp.Data.Sex,
		Age:      resp.Data.Age,
	}
}
