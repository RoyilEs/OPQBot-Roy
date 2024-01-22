package group

import (
	"fmt"
	"obqbot/global"
	"obqbot/models"
)

func InitMeme() {
	friendTag := models.FriendTag{
		Name: "rtwyzz",
		TagsData: models.StringArray{
			"我操死你的吗rtwyzz",
			"怂逼,不敢对嘴",
		},
	}
	global.DB.Create(&friendTag)

	global.DB.Where("name = ?", "rwtyzz").First(&friendTag)
	fmt.Println(friendTag.TagsData[0])
}
