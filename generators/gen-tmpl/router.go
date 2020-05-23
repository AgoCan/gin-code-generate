package gen_tmpl

var RouterContent = `package routers
import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/middleware"
)
// SetupRouter 路由路口
func SetupRouter() *gin.Engine{
	router := gin.Default()
	router.Use(middleware.LogMiddleware())
	return router
}`