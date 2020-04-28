package httpSever

import (
	"gateway/config"
	"gateway/tool/log"
	"go.uber.org/zap"
)

//启动http 路由服务
func StartHttpServer() {
	httpServer := NewHttpServer(config.GetConfig().HttpServer.Addr)
	go httpServer.Start()
	log.Instance().Info("info", zap.Any("httpServer", "server start..启动成功 端口:"+config.GetConfig().HttpServer.Addr))

}
