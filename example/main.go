package main

import (
	Lagrange "Lagrange-SDK"
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"Lagrange-SDK/example/listens"
	"Lagrange-SDK/utils/http"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

func main() {
	core, err := Lagrange.NewCore("8.141.1.249:8081")
	if err != nil {
		return
	}
	core.On(events.EventGroupMsg, func(client *websocket.Conn, event events.IEvent) {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()
		log.Info(text)

		if text[0] == "test" {
			apiBuilder.NewApi().GetGroupMemberInfo().ToGroupIDAndUserID(groupMsg.GetGroupID(), groupMsg.GetUserID()).Do(client)
			groupMember := event.ParseGroupMemberInfo()
			log.Info(groupMember.GetNickName())
			apiBuilder.NewApi().SendReply(event.GetMessageID()).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("回复测试").Do(client)
		}

		if text[0] == "信息体获取" {
			apiBuilder.NewApi().SendGroupMsg(groupMsg.GetGroupID()).
				TextMsg(fmt.Sprintf("%v", event.GetEventMessageStruct())).Do(client)
		}

		if text[0] == "image" {
			httpClient := http.NewHTTPClient("https://api.likepoems.com/img/pc/")
			doGet, _ := httpClient.DoGet("", nil)

			apiBuilder.NewApi().SendGroupMsg(groupMsg.GetGroupID()).
				ImgBase64Msg(base64.StdEncoding.EncodeToString(doGet)).Do(client)
		}
	})

	core.On(events.EventGroupMsg, listens.BanKick)
	core.On(events.EventGroupMsg, listens.PixivImg)
	core.On(events.EventGroupMsg, listens.ArknightsImg)
	core.On(events.EventGroupMsg, listens.MingImg)

	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
