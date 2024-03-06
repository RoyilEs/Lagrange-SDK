package pixiv

import "net/url"

var PixivUrl = "https://api.lolicon.app/setu/v2"

type ISetQuery interface {
	SetSize(size string) ISetQuery
	SetTag(tag string) ISetQuery
	DoQuery() (*Pixiv, error)
}

type IGetQuery interface {
	GetSizeQuery() string
	GetTagQuery() string
}

func (p *Pixiv) GetSizeQuery() string {
	return p.size
}
func (p *Pixiv) GetTagQuery() string {
	return p.tag
}

func (p *Pixiv) Set() ISetQuery {
	p.size = "regular"
	p.tag = "%E6%98%8E%E6%97%A5%E6%96%B9%E8%88%9F"
	return p
}

func (p *Pixiv) SetSize(size string) ISetQuery {
	p.size = size
	return p
}

func (p *Pixiv) SetTag(tag string) ISetQuery {
	p.tag = url.QueryEscape(tag)
	return p
}

func (p *Pixiv) DoQuery() (*Pixiv, error) {
	return p, nil
}
