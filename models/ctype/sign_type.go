package ctype

import "encoding/json"

type SignType int

const (
	SignOk SignType = 1 //已签到
	SignNo SignType = 0 //未签到
)

func (r SignType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r SignType) String() interface{} {
	var str string
	switch r {
	case SignOk:
		str = "已签到"
	case SignNo:
		str = "未签到"
	default:
		str = "未知"
	}
	return str
}
