package main

import (
	"Lagrange-SDK/events"
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"strconv"
)

func ListenGroup(_ context.Context, event events.IEvent) {
	if event.GetMessageType() != events.EventGroupMsg {
		return
	}
	groupMsg := event.ParseGroupMsg() //群消息
	groupID := "群:" + strconv.FormatInt(groupMsg.GetGroupID(), 10)
	user := fmt.Sprintf("成员: %s (%d)", groupID, event.GetUserID())
	log.Info(user)

	log.Info(fmt.Println(fmt.Sprintf("信息: %s", groupMsg.ParseTextMsg())))
}
