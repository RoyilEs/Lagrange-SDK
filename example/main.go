package main

import (
	Lagrange "Lagrange-SDK"
	"Lagrange-SDK/apiBuilder"
	"Lagrange-SDK/events"
	"Lagrange-SDK/global"
	"Lagrange-SDK/models/pixiv2"
	http2 "Lagrange-SDK/utils/http"
	"Lagrange-SDK/utils/image"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/RoyilEs/asiatz"
	"github.com/charmbracelet/log"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	core, err := Lagrange.NewCore("ws://8.152.216.18:3001", 3392313023)
	if err != nil {
		return
	}
	core.On(events.EventGroupMsg, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.Message().ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()
		log.Info(text)

		if text[0] == "test" {
			groupMember := event.Message().ParseGroupMsg()
			log.Info(groupMember.GetNickName())
			apiBuilder.New(global.BotUrl).
				SendGroupMsg(groupMsg.GetGroupID()).TextMsg("vivo50").Do(ctx)
		}

		if text[0] == "信息体获取" {
			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
				TextMsg(fmt.Sprintf("%v", event.Message().GetEventMessageStruct())).Do(ctx)
		}

		if text[0] == "image" {
			httpClient := http2.NewHTTPClient("https://api.likepoems.com/img/pc/")
			doGet, _ := httpClient.DoGet("", nil)

			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
				ImgBase64Msg(base64.StdEncoding.EncodeToString(doGet)).Do(ctx)
		}

		if text[0] == "asiatc" {
			utcTime, err := asiatz.ShanghaiToUTC("10:00")
			if err != nil {
				log.Error(err)
			}
			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
				TextMsg(utcTime).Do(ctx)
		}

	})
	core.On(events.EventGroupMsg, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.Message().ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()
		log.Info(text)
		parts := strings.Split(text[0], " ")
		if parts[0] == "amd" {
			url := "http://dkdj.qlit.edu.cn:9901/api/getSydl"
			method := "POST"

			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			_ = writer.WriteField("loudong", parts[1]+"号公寓照明")
			_ = writer.WriteField("room", parts[2]+"照明")
			_ = writer.WriteField("xiaoqu", "济南校区")
			err := writer.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			client := &http.Client{}
			req, err := http.NewRequest(method, url, payload)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
				TextMsg(string(body)).Do(ctx)
		}
	})
	core.On(events.EventGroupMsg, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.Message().ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetText()
		parts := strings.Split(text[0], " ")
		if parts[0] == "pixiv" {
			client := &http.Client{
				Timeout: time.Second * 30, // 设置请求超时时间为30秒
			}
			log.Info("pixiv")
			if len(parts) > 1 {
				pid, title, url, err := pixiv2.GetPixivPidTitleUrl(client, groupMsg.GetGroupID(), parts[1], 0, ctx)
				if err != nil {
					return
				}
				log.Info(pid, title, url)
				log.Info("通过base64压缩中")
				encodeToBase64, err := image.CompressQualityAndEncodeToBase64ByUrl(url, 50)
				if err != nil {
					apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg(err.Error()).Do(ctx)
					apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("压缩失败").Do(ctx)
					apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("原图链接: " + url).Do(ctx)
					log.Error(err)
					return
				}
				apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
					TextMsg(fmt.Sprintf("pid: %v, title: %v, url: %v", pid, title, url)).ImgBase64Msg(encodeToBase64).Do(ctx)
				log.Info("已发送")
			} else {
				pid, title, url, err := pixiv2.GetPixivPidTitleUrl(client, groupMsg.GetGroupID(), "", 0, ctx)
				if err != nil {
					return
				}
				log.Info(pid, title, url)
				log.Info("通过base64压缩中")
				encodeToBase64, err := image.CompressQualityAndEncodeToBase64ByUrl(url, 50)
				if err != nil {
					apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg(err.Error()).Do(ctx)
					apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("压缩失败").Do(ctx)
					apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).TextMsg("原图链接: " + url).Do(ctx)
					log.Error(err)
					return
				}
				apiBuilder.New(global.BotUrl).SendGroupMsg(groupMsg.GetGroupID()).
					TextMsg(fmt.Sprintf("pid: %v, title: %v, url: %v", pid, title, url)).ImgBase64Msg(encodeToBase64).Do(ctx)
				log.Info("已发送")
			}
		}

		if parts[0] == "pixiv18" {
			client := &http.Client{
				Timeout: time.Second * 30, // 设置请求超时时间为30秒
			}
			log.Info("pixiv")
			if len(parts) > 1 {
				pid, title, url, err := pixiv2.GetPixivPidTitleUrl(client, groupMsg.GetGroupID(), parts[1], 1, ctx)
				if err != nil {
					return
				}
				log.Info(pid, title, url)
				log.Info("通过base64压缩中")
				encodeToBase64, err := image.CompressQualityAndEncodeToBase64ByUrl(url, 50)
				if err != nil {
					apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).TextMsg(err.Error()).Do(ctx)
					apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).TextMsg("压缩失败").Do(ctx)
					apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).TextMsg("原图链接: " + url).Do(ctx)
					log.Error(err)
					return
				}
				apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).
					TextMsg(fmt.Sprintf("pid: %v, title: %v, url: %v", pid, title, url)).ImgBase64Msg(encodeToBase64).Do(ctx)
				log.Info("已发送")
			} else {
				pid, title, url, err := pixiv2.GetPixivPidTitleUrl(client, groupMsg.GetGroupID(), "", 1, ctx)
				if err != nil {
					return
				}
				log.Info(pid, title, url)
				log.Info("通过base64压缩中")
				encodeToBase64, err := image.CompressQualityAndEncodeToBase64ByUrl(url, 50)
				if err != nil {
					apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).TextMsg(err.Error()).Do(ctx)
					apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).TextMsg("压缩失败").Do(ctx)
					apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).TextMsg("原图链接: " + url).Do(ctx)
					log.Error(err)
					return
				}
				apiBuilder.New(global.BotUrl).SendPrivateMsg(groupMsg.GetSenderUserID()).
					TextMsg(fmt.Sprintf("pid: %v, title: %v, url: %v", pid, title, url)).ImgBase64Msg(encodeToBase64).Do(ctx)
				log.Info("已发送")
			}
		}
	})
	core.On(events.EventGroupMsg, func(ctx context.Context, event events.IEvent) {
		msg := event.Message().ParseGroupMsg()
		text := msg.ParseTextMsg().GetText()
		fmt.Println(msg.ParseTextMsg().GetTextData())
		qq := msg.ParseTextMsg().GetAtQQ()
		fmt.Println(qq)
		if text[0] == "poke " {
			atoi, _ := strconv.Atoi(qq[0])
			apiBuilder.New(global.BotUrl).Poke().SetGroupPoke(msg.GetGroupID(), int64(atoi)).Do(ctx)
		}
	})
	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
