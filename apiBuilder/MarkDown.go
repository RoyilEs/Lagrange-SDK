package apiBuilder

import (
	"Lagrange-SDK/message"
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"net/url"
	"strconv"
	"strings"
)

type IMarkDownBuild interface {
	Set(MarkDown) IMarkDownBuild
	DoForWardID(ctx context.Context) (string, error)
}

type MarkDown struct {
	Str string
}

func NewMarkDown() *MarkDown {
	return &MarkDown{
		Str: "",
	}
}

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

func (b *Builder) DoForWardID(ctx context.Context) (string, error) {
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

func (b *Builder) Set(markDown MarkDown) IMarkDownBuild {
	for _, messageStruct := range b.Params.Messages {
		m := &MarkDownBuild{
			Type: message.MARKDOWN,
			Data: struct {
				Content string `json:"content,omitempty"`
			}{Content: markDown.Str},
		}
		i := append(messageStruct.Data.Content, m)
		messageStruct.Data.Content = i
	}
	return b
}

func (r *MarkDown) MqqApi(content string) *MarkDown {
	str := fmt.Sprintf("\\n[%s](mqqapi://aio/inlinecmd?command=%s&reply=false&enter=false)", content, url.PathEscape(content))
	r.Str += str
	return r
}

func (r *MarkDown) MqqApiAuto(content string) *MarkDown {
	str := fmt.Sprintf("\\n[%s](mqqapi://aio/inlinecmd?command=%s&reply=false&enter=true)", content, url.PathEscape(content))
	r.Str += str
	return r
}

// MqqApiAt 需配合 SendMarkdownAtMsg 和 SendMarkdownAndKeyboardAtMsg 使用，否则没有效果
func (r *MarkDown) MqqApiAt(nickname string, tinyId uint64) *MarkDown {
	str := fmt.Sprintf("\\n[@%s](mqqapi://markdown/mention?at_type=1&at_tinyid=%v)", nickname, tinyId)
	r.Str += str
	return r
}

// MqqApiAtToProfile 需配合 SendMarkdownAtMsg 和 SendMarkdownAndKeyboardAtMsg 使用，否则没有效果
func (r *MarkDown) MqqApiAtToProfile(nickname string, tinyId uint64) *MarkDown {
	str := fmt.Sprintf("\\n[@%s](mqqapi://card/show_pslcard?src_type=internal&version=1&uin=%v&crad_type=friend&source=qrcode)", nickname, tinyId)
	r.Str = str
	return r
}

func (r *MarkDown) Url(name, webUrl string) *MarkDown {
	str := fmt.Sprintf("\\n[🔗%s](%s)", name, webUrl)
	r.Str += str
	return r
}

func (r *MarkDown) H1(content string) *MarkDown {
	str := fmt.Sprintf("\\n# %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) H2(content string) *MarkDown {
	str := fmt.Sprintf("\\n## %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) H3(content string) *MarkDown {
	str := fmt.Sprintf("\\n### %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) DeleteLine(content string) *MarkDown {
	str := fmt.Sprintf("~~%s~~", content)
	r.Str += str
	return r
}

func (r *MarkDown) Bold(content string) *MarkDown {
	str := fmt.Sprintf("**%s** ", content)
	r.Str += str
	return r
}

func (r *MarkDown) Italic(content string) *MarkDown {
	str := fmt.Sprintf("*%s* ", content)
	r.Str += str
	return r
}

func (r *MarkDown) ItalicBold(content string) *MarkDown {
	str := fmt.Sprintf("***%s*** ", content)
	r.Str += str
	return r
}

func (r *MarkDown) BlockReference(content string) *MarkDown {
	str := fmt.Sprintf("\\n> %s\\n", content)
	r.Str += str
	return r
}

func (r *MarkDown) Image(text, url string, width, height int) *MarkDown {
	str := fmt.Sprintf("\\n![%s #%vpx #%vpx](%s)", text, width, height, url)
	r.Str += str
	return r
}

func (r *MarkDown) DividerLine() *MarkDown {
	str := fmt.Sprintln("\\n ---\\n")
	r.Str += str
	r.Str = strings.TrimSpace(r.Str)
	return r
}

func (r *MarkDown) Text(content string) *MarkDown {
	r.Str += content
	return r
}

func (r *MarkDown) NewLine() *MarkDown {
	r.Str += "\\n"
	return r
}

func (r *MarkDown) Code(content string) *MarkDown {
	c := strings.ReplaceAll(strings.ReplaceAll(content, "\t", "\\t"), "\n", "\\n")
	str := fmt.Sprintf("\\n```\\n%s\\n```\\n", c)
	r.Str += str
	return r
}
