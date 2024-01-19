package global

import (
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"obqbot/config"
)

// 全局变量
var (
	Config        *config.Config
	DB            *gorm.DB
	Log           *logrus.Logger
	MySqlLog      logger.Interface
	ApiBuilderNew = apiBuilder.New(OBQBotUrl, BotQQ)

	// GroupUids 有权限的 群组
	GroupUids = []int64{12313}

	// AdminUids 管理员
	AdminUids = []int64{2839182980, 3392313023}
)

const (
	OBQBotUrl = "http://8.141.1.249:9000"
	BotQQ     = 3392313023
)
