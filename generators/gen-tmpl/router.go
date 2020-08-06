package tmpl

// RouterContent router
var RouterContent = `package routers

import (
	"{{ .ProjectName }}/middleware/log"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化gin入口，路由信息
func SetupRouter() *gin.Engine{
	router := gin.New()
	if err := log.InitLogger(); err != nil {
		panic(err)
	}
	router.Use(log.GinLogger(log.Logger),
		log.GinRecovery(log.Logger, true))

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	return router
}
`
