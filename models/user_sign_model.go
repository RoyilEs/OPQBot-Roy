package models

import (
	"gorm.io/gorm"
	"obqbot/models/ctype"
)

// UserSign 用户签到
type UserSign struct {
	gorm.Model
	NickName string         `json:"nickname" gorm:"type:varchar(255);not null"` //昵称
	Uin      int64          `json:"uin" gorm:"type:bigint(20);not null"`        //uin
	GroupUin int64          `json:"group_uin" gorm:"type:bigint(20);not null"`  //群uin
	SignType ctype.SignType `json:"sign_type" gorm:"type:tinyint(1);not null"`  //签到类型
	Point    int64          `json:"point" gorm:"type:bigint(20);not null"`      //积分
}
