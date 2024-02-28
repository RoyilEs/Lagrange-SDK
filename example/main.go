package main

import (
	Lagrange "Lagrange-SDK"
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"Lagrange-SDK/utils/http"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

func main() {
	core, err := Lagrange.NewCore("127.0.0.1:8080")
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
			apiBuilder.NewApi().SendGroupMsg().ToGroupID(groupMsg.GetGroupID()).
				TextMsg(fmt.Sprintf("%d--%d\n%s\n%s",
					groupMember.GetGroupID(), groupMember.GetUserID(),
					groupMember.GetNickName(), groupMember.GetLevel()) + " 成功").Do(client)
		}
		if text[0] == "image" {
			httpClient := http.NewHTTPClient("https://api.likepoems.com/img/pc/")
			doGet, _ := httpClient.DoGet("", nil)

			apiBuilder.NewApi().SendGroupMsg().ToGroupID(groupMsg.GetGroupID()).
				ImgBase64Msg(base64.StdEncoding.EncodeToString(doGet)).Do(client)
		}
	})

	core.On(events.EventSetAdmin, func(client *websocket.Conn, event events.IEvent) {
		set := event.ParseSet()
		log.Info(set)
		apiBuilder.NewApi().SendGroupMsg().ToGroupID(set.GetGroupID()).
			TextMsg(fmt.Sprintf("%d 成功被设置为管理员", set.GetUserID())).Do(client)
	})

	core.On(events.EventUnSetAdmin, func(client *websocket.Conn, event events.IEvent) {
		set := event.ParseSet()
		log.Info(set)
		apiBuilder.NewApi().GetGroupMemberInfo().ToGroupIDAndUserID(set.GetGroupID(), set.GetUserID()).Do(client)
		groupMember := event.ParseGroupMemberInfo()
		log.Info(groupMember.GetNickName())
		apiBuilder.NewApi().SendGroupMsg().ToGroupID(set.GetGroupID()).
			TextMsg(fmt.Sprintf("%d--%d\n%s\n%s",
				groupMember.GetGroupID(), groupMember.GetUserID(),
				groupMember.GetNickName(), groupMember.GetLevel()) + " 成功").Do(client)

		apiBuilder.NewApi().SendGroupMsg().ToGroupID(set.GetGroupID()).
			TextMsg(fmt.Sprintf("%d 成功被取消为管理员", set.GetUserID())).Do(client)
	})

	core.On(events.EventKick, func(client *websocket.Conn, event events.IEvent) {
		kick := event.ParseKick()
		log.Info(kick)
		apiBuilder.NewApi().SendGroupMsg().ToGroupID(kick.GetGroupID()).
			TextMsg(fmt.Sprintf("%d 成功被踢出群聊", kick.GetUserID())).Do(client)
	})

	core.On(events.EventInvite, func(client *websocket.Conn, event events.IEvent) {
		kick := event.ParseKick()
		log.Info(kick)
		apiBuilder.NewApi().SendGroupMsg().ToGroupID(kick.GetGroupID()).
			TextMsg(fmt.Sprintf("%d 成功加入群聊", kick.GetUserID())).Do(client)
	})

	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
