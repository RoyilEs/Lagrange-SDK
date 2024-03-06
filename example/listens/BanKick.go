package listens

import (
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"strconv"
	"strings"
)

func BanKick(client *websocket.Conn, event events.IEvent) {
	if string(events.EventGroupMsg) != event.GetMessageType() {
		return
	}
	groupMsg := event.ParseGroupMsg()
	text := groupMsg.ParseTextMsg().GetText()

	if text[0] == "mute " {
		atoi, err := strconv.Atoi(strings.TrimSpace(text[2]))
		if err != nil {
			apiBuilder.NewApi().SendGroupMsg(groupMsg.GetGroupID()).TextMsg("请输入数字1").Do(client)
			log.Info(err)
			return
		}

		atQQ := groupMsg.ParseTextMsg().GetAtQQ()
		i, err := strconv.ParseInt(atQQ[0], 10, 64)
		if err != nil {
			apiBuilder.NewApi().SendGroupMsg(groupMsg.GetGroupID()).TextMsg("请输入数字2").Do(client)
			return
		}
		apiBuilder.NewApi().SetGroupBan().ToGroupIDAndMuteUserID(groupMsg.GetGroupID(), i).Duration(atoi).Do(client)
	}

	if text[0] == "kick " {
		fmt.Println(text)
		atQQ := groupMsg.ParseTextMsg().GetAtQQ()
		i, err := strconv.ParseInt(atQQ[0], 10, 64)
		if err != nil {
			apiBuilder.NewApi().SendGroupMsg(groupMsg.GetGroupID()).TextMsg("请输入数字2").Do(client)
			return
		}
		apiBuilder.NewApi().SetGroupKick().ToGroupIDAndKickUserID(groupMsg.GetGroupID(), i).Do(client)
	}
}
