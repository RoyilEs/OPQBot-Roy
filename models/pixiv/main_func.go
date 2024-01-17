package pixiv

type IPixiv interface {
	Do(url string, pixiv *Pixiv) (*PixivResponse, error)
	Set() ISetQuery
}
