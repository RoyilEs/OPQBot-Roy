package main

import (
	"obqbot/core"
	"obqbot/global"
	"obqbot/listener/group"
)

func main() {
	core.InitConf()
	global.DB = core.InitGorm()

	group.InitMeme()
}
