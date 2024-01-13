package global

import (
	"github.com/sirupsen/logrus"
	"obqbot/config"
)

// 全局变量
var (
	Config *config.Config
	Log    *logrus.Logger
)

const (
	OBQBotUrl = "http://localhost:8086"
	BotQQ     = 3392313023
)
