package pixiv

type IPixiv interface {
	Set() ISetQuery
	Do(url string, pixiv *Pixiv) (*PixivResponse, error)
}
