package apiBuilder

import (
	"Lagrange-SDK/message"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"image"
	"strconv"
)

type IMarkDownBuild interface {
	SetImage(img string) IMarkDownBuild
	SetGo(goLagrange string) IMarkDownBuild
	DoBuild(ctx context.Context) (string, error)
}

var (
	text, content string
)

func (b *Builder) MarkDownBuild(name string, uin int64) IMarkDownBuild {
	cmd := SendGroupForwardMsg
	b.action = cmd
	i := append(b.Params.Messages, &MessageStruct{
		Type: message.NODE,
		Data: DataMessage{
			Name: name,
			Uin:  strconv.FormatInt(uin, 10),
		},
	})
	b.Params.Messages = i
	return b
}

func (b *Builder) SetImage(img string) IMarkDownBuild {
	for _, messageStruct := range b.Params.Messages {
		imgBytes, err := base64.StdEncoding.DecodeString(img)
		var (
			height int
			width  int
		)
		if err == nil {
			h, w := GetImageHW(imgBytes)
			height = h
			width = w
		} else {
			log.Debug(err)
		}
		if height == 0 || width == 0 {
			height = 1920
			width = 1080
		}
		text += fmt.Sprintf("![text #%d #%d](%s)", height, width, img)
		content = "{\"content\":\" " + text + " \" \n}"
		m := &MarkDownBuild{
			Type: message.MARKDOWN,
			Data: struct {
				Content string `json:"content,omitempty"`
			}{Content: content},
		}
		i := append(messageStruct.Data.Content, m)
		messageStruct.Data.Content = i
		text = ""
	}
	return b
}

func (b *Builder) SetGo(goLagrange string) IMarkDownBuild {
	for _, messageStruct := range b.Params.Messages {
		text += fmt.Sprintf("```go\n%s\n```", goLagrange)
		content = "{\"content\":\" " + text + " \" \n}"
		m := &MarkDownBuild{
			Type: message.MARKDOWN,
			Data: struct {
				Content string `json:"content,omitempty"`
			}{Content: content},
		}
		i := append(messageStruct.Data.Content, m)
		messageStruct.Data.Content = i
		text = ""
	}
	return b
}

func (b *Builder) DoBuild(ctx context.Context) (string, error) {
	resp, err := b.DoAndResponse(ctx)
	if err != nil {
		return "nil", err
	}
	log.Debug(string(resp.GetOrigin()))
	if !resp.Ok() {
		return "", errors.New(resp.StatusMsg())
	}

	if err != nil {
		return "", err
	}
	return resp.response.Data.ForwardID, nil
}

func GetImageHW(pic []byte) (height, width int) {
	img, _, _ := image.Decode(bytes.NewBuffer(pic))
	return img.Bounds().Dy(), img.Bounds().Dx()
}
