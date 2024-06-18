package listens

import (
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"Lagrange-SDK/global"
	"context"
	"strconv"
	"strings"
)

func GetUser(ctx context.Context, event events.IEvent) {
	if string(events.EventGroupMsg) != event.GetMessageType() {
		return
	}
	groupMsg := event.ParseGroupMsg()
	text := groupMsg.ParseTextMsg().GetText()
	split := strings.Split(text[0], " ")

	if split[0] == "getUser" {
		i, _ := strconv.ParseInt(split[1], 10, 64)
		getStrangerInfo := apiBuilder.New(global.BotUrl).GetStrangerInfo(ctx, i)
		apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
			TextMsg(getStrangerInfo.NickName + "\t" + strconv.FormatInt(getStrangerInfo.UserID, 10)).Do(ctx)
	}

}
