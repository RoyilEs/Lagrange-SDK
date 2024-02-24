# LagrangeBot Golang SDK 🎉
欢迎 Star 👍

## 使用说明

```go
package main

import (
	Lagrange "Lagrange-SDK"
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"context"
)

core, err := Lagrange.NewCore("8.141.1.249:8081")
	if err != nil {
		return
	}
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
