package pixiv

import (
	"encoding/json"
	"fmt"
	Ok3Http "obqbot/utils"
)

var pixivUrl = "https://api.lolicon.app/setu/v2"

type ISetQuery interface {
	SetSize(size string) ISetQuery
	SetTag(tag string) ISetQuery
	DoPixiv() (*PixivResponse, error)
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
	p.tag = tag
	return p
}

func (p *Pixiv) DoPixiv() (*PixivResponse, error) {
	var (
		size = p.GetSizeQuery()
		tag  = p.GetTagQuery()
	)
	pixivUrl = pixivUrl + "?size=" + size + "&tag=" + tag + "&r18=0"
	fmt.Println(pixivUrl)
	fmt.Println(p.GetTagQuery())
	body, err := Ok3Http.NewHTTPClient(pixivUrl).DoGet("", nil)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	pixivResponse := PixivResponse{}
	err = json.Unmarshal(body, &pixivResponse)
	if err != nil {
		return nil, err
	}
	return &pixivResponse, nil
}
