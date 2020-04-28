package httpSever

import (
	"gateway/auth"
	"gateway/tool/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	r.Use(setCROSOptions)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "sdfsdfsdfsdf")
	})

	private := r.Group("/v1", auth.CheckToken())
	{
		private.GET("/user", func(c *gin.Context) {

			c.String(http.StatusOK, "这是所有用户列表")

		})
		private.GET("/user/:user_id", func(c *gin.Context) {

			c.String(http.StatusOK, "这是单个用户列表")

		}) //测试接口

	}
	err := r.Run(server.addr)
	if err != nil {
		log.Instance().Panic("panic", zap.Any("httpServer-err", err))
	}

}

func setCROSOptions(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}

}
