package listens

import (
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"github.com/gorilla/websocket"
)

func Hello(client *websocket.Conn, event events.IEvent) {
	groupMsg := event.ParseGroupMsg()
	if groupMsg.ParseTextMsg().GetText()[0] == "hello" {
		apiBuilder.New().SendGroupMsg().ToGroupID(groupMsg.GetGroupID()).TextMsg("world").Do(client)
	}
}
