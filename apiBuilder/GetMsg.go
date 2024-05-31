package apiBuilder

import (
	"context"
	"github.com/charmbracelet/log"
)

type MsgStruct struct {
	Time        int64  `json:"time"`
	MessageType string `json:"message_type"`
	MessageID   int64  `json:"message_id"`
	RealID      int64  `json:"real_id"`
	Sender      struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
	} `json:"sender"`
	Data []*MessageStruct
}

func (b *Builder) GetMsg(ctx context.Context, msgID int64) MsgStruct {
	b.action = GetMsg
	b.Params.MessageID = msgID
	resp, err := b.DoResponse(ctx)
	if err != nil {
		log.Error(err)
		return MsgStruct{}
	}

	var data []*MessageStruct
	for _, v := range resp.Data.Message {
		data = append(data, &MessageStruct{
			Type: v.Type,
			Data: DataMessage{
				File: v.Data.File,
				Text: v.Data.Text,
				Data: v.Data.Data,
				Url:  v.Data.Url,
				QQ:   v.Data.QQ,
				ID:   v.Data.ID,
			},
		})
	}

	return MsgStruct{
		Time:        resp.Data.Time,
		MessageType: resp.Data.MessageType,
		MessageID:   resp.Data.MessageID,
		RealID:      resp.Data.RealID,
		Sender: struct {
			UserID   int64  `json:"user_id"`
			Nickname string `json:"nickname"`
			Sex      string `json:"sex"`
		}{
			UserID:   resp.Data.Sender.UserID,
			Nickname: resp.Data.NickName,
			Sex:      resp.Data.Sex,
		},
		Data: data,
	}
}
