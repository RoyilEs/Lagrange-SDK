package listens

import (
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"Lagrange-SDK/global"
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"strconv"
	"strings"
)

func BanKick(ctx context.Context, event events.IEvent) {
	if string(events.EventGroupMsg) != event.GetMessageType() {
		return
	}
	groupMsg := event.ParseGroupMsg()
	text := groupMsg.ParseTextMsg().GetText()

	if text[0] == "mute " {
		atoi, err := strconv.Atoi(strings.TrimSpace(text[2]))
		if err != nil {
			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("请输入数字1").Do(ctx)
			log.Info(err)
			return
		}

		atQQ := groupMsg.ParseTextMsg().GetAtQQ()
		i, err := strconv.ParseInt(atQQ[0], 10, 64)
		if err != nil {
			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("请输入数字2").Do(ctx)
			return
		}
		apiBuilder.New(global.BotUrl).SetGroupBan().ToGroupIDAndMuteUserID(groupMsg.GetGroupID(), i).Duration(atoi).Do(ctx)
	}

	if text[0] == "kick " {
		fmt.Println(text)
		atQQ := groupMsg.ParseTextMsg().GetAtQQ()
		i, err := strconv.ParseInt(atQQ[0], 10, 64)
		if err != nil {
			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("请输入数字2").Do(ctx)
			return
		}
		apiBuilder.New(global.BotUrl).SetGroupKick().ToGroupIDAndKickUserID(groupMsg.GetGroupID(), i).Do(ctx)
	}
}
