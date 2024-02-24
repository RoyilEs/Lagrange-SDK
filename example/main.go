package main

import (
	Lagrange "Lagrange-SDK"
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"context"
	"fmt"
)

func main() {
	core, err := Lagrange.NewCore("127.0.0.1:8080")
	if err != nil {
		return
	}
	core.On(events.EventPrivateMsg, func(ctx context.Context, event events.IEvent) {
		fmt.Println(event.GetMessageType())
	})
	core.On(events.EventGroupMsg, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.ParseGroupMsg()
		if groupMsg.ParseTextMsg().GetText()[0] == "test" {
			apiBuilder.New().SendGroupMsg().ToGroupID(groupMsg.GetGroupID()).TextMsg("测试").Do()
		}

	})
	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
