package main

import (
	"fmt"
	"obqbot/models/pixiv"
)

func main() {
	data := pixiv.NewPixiv().GetData()
	fmt.Println(data)
}
