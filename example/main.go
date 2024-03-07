package main

import (
	Lagrange "Lagrange-SDK"
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"Lagrange-SDK/global"
	"Lagrange-SDK/utils/http"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/charmbracelet/log"
)

func main() {
	core, err := Lagrange.NewCore(global.BotWs)
	if err != nil {
		return
	}
	core.On(events.EventGroupMsg, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()
		log.Info(text)

		if text[0] == "test" {
			apiBuilder.New(global.BotUrl).GetGroupMemberInfo().ToGroupIDAndUserID(groupMsg.GetGroupID(), groupMsg.GetUserID()).Do(ctx)
			groupMember := event.ParseGroupMemberInfo()
			log.Info(groupMember.GetNickName())
			apiBuilder.New(global.BotUrl).SendReply(event.GetMessageID()).
				SendGroupMsg(groupMsg.GetGroupID()).TextMsg("回复测试").Do(ctx)
		}

		if text[0] == "信息体获取" {
			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
				TextMsg(fmt.Sprintf("%v", event.GetEventMessageStruct())).Do(ctx)
		}

		if text[0] == "image" {
			httpClient := http.NewHTTPClient("https://api.likepoems.com/img/pc/")
			doGet, _ := httpClient.DoGet("", nil)

			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
				ImgBase64Msg(base64.StdEncoding.EncodeToString(doGet)).Do(ctx)
		}
	})

	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
