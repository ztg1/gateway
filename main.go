package main

import (
	"fmt"
	"gateway/config"
	"gateway/tool/log"
	"go.uber.org/zap"
)

func main() {
	log.Init("gataeway-start")
	data := config.GetConfig()
	fmt.Println("dataConfig=", data)

	log.Instance().Info("info", zap.Any("start", "server-start......"))
}
