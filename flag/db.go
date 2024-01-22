package flag

import (
	"fmt"
	"obqbot/global"
	"obqbot/models"
)

func MakeMigrations() {
	var err error
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserSign{},
			&models.FriendTag{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败", err)
		return
	}
	fmt.Println("MakeMigrations OPQbot-Roy [ ok ]")
}
