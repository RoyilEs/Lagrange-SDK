package pixiv

import (
	"Lagrange-SDK/utils/http"
	"encoding/base64"
	"github.com/charmbracelet/log"
	jsoniter "github.com/json-iterator/go"

	"strings"
)

type Pixiv struct {
	size string
	tag  string
}

type PixivResponse struct {
	Error string `json:"error"`
	Data  []Data `json:"data"`
}

type Data struct {
	Pid        int64    `json:"pid"`
	P          int64    `json:"p"`
	Uid        int64    `json:"uid"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	R18        bool     `json:"r18"`
	Width      int64    `json:"width"`
	Height     int64    `json:"height"`
	Tags       []string `json:"tags"`
	Ext        string   `json:"ext"`
	AiType     int64    `json:"aiType"`
	UploadDate int64    `json:"uploadDate"`
	Urls       Urls     `json:"urls"`
}

type Urls struct {
	Regular string `json:"regular"`
}

type IPixivData interface {
	GetDataPid() int64
	GetDataP() int64
	GetDataUid() int64
	GetDataTitle() string
	GetDataAuthor() string
	GetDataR18() bool
	GetDataWidth() int64
	GetDataHeight() int64
	GetDataTags() []string
	GetDataExt() string
	GetDataAiType() int64
	GetDataUploadDate() int64
	GetDataUrls() Urls
}

type IData interface {
	GetData() []Data
}

type IDataUrls interface {
	GetSize() string
}

func (u Urls) GetSize() string {
	return u.Regular
}

func (p *PixivResponse) GetData() []Data {
	return p.Data
}

func (p *Data) GetDataPid() int64 {
	return p.Pid
}
func (p *Data) GetDataP() int64 {
	return p.P
}
func (p *Data) GetDataUid() int64 {
	return p.Uid
}

func (p *Data) GetDataTitle() string {
	return p.Title
}
func (p *Data) GetDataAuthor() string {
	return p.Author
}
func (p *Data) GetDataR18() bool {
	return p.R18
}
func (p *Data) GetDataWidth() int64 {
	return p.Width
}
func (p *Data) GetDataHeight() int64 {
	return p.Height
}
func (p *Data) GetDataTags() []string {
	return p.Tags
}
func (p *Data) GetDataExt() string {
	return p.Ext
}
func (p *Data) GetDataAiType() int64 {
	return p.AiType
}
func (p *Data) GetDataUploadDate() int64 {
	return p.UploadDate
}
func (p *Data) GetDataUrls() Urls {
	return p.Urls
}

func ModifyPixivImageUrl(originalUrl string) (modifiedUrl string) {
	// 分离出基础URL和尺寸部分
	parts := strings.Split(originalUrl, "/")
	// 查找包含"_master"的部分
	squareIndex := -1
	for i, part := range parts {
		if strings.Contains(part, "_master") {
			squareIndex = i
			break
		}
	}

	// 如果找到了"_master"，则替换为"_square"
	if squareIndex >= 0 {
		modifiedPart := strings.Replace(parts[squareIndex], "_master", "_square", 1)
		parts[squareIndex] = modifiedPart
		modifiedUrl = strings.Join(parts, "/")
	} else {
		// 如果没有找到"_master"，则返回原URL
		modifiedUrl = originalUrl
	}

	split := strings.Split(modifiedUrl, "/")
	split = insertSlice(split, 3, "c")
	split = insertSlice(split, 4, "250x250_80_a2")
	//slice := deleteSlice(split, "c")
	//modifiedUrl = strings.Join(deleteSlice(slice, "250x250_80_a2"), "/")

	return strings.Join(split, "/")
}

func insertSlice(s []string, index int, elem string) []string {
	s = append(s, elem)
	copy(s[index+1:], s[index:])
	s[index] = elem
	return s
}

// deleteSlice 删除指定元素。
func deleteSlice(s []string, elem string) []string {
	r := make([]string, 0, len(s))
	for _, v := range s {
		if v != elem {
			r = append(r, v)
		}
	}
	return r
}

func NewPixiv() IPixiv {
	return &Pixiv{}
}

func (p *Pixiv) Do(url string, pixiv *Pixiv) (*PixivResponse, error) {
	url = url + "?size=" + pixiv.GetSizeQuery() + "&tag=" + pixiv.GetTagQuery()
	s := http.NewHTTPClient(url)
	body, err := s.DoGet("", nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var pixivResponse PixivResponse

	err = jsoniter.NewDecoder(strings.NewReader(string(body))).Decode(&pixivResponse)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &pixivResponse, nil
}

func (p *Data) UrlToBase64(url string) (base64buf string) {

	s := http.NewHTTPClient(url)
	body, err := s.DoGet("", nil)
	if err != nil {
		log.Error(err)
		return ""
	}
	base64buf = base64.StdEncoding.EncodeToString(body)
	return base64buf
}
