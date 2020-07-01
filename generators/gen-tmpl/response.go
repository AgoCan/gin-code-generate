package tmpl

// ResponseContent 回调使用
var ResponseContent = `package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Data 回调时的固定地址
type Data struct {
	Code    int         ` + "`json:" + `"code"` + "`" + `
	Message string      ` + "`json:" + `"message"` + "`" + `
	Data    interface{} ` + "`json:" + `"data"` + "`" + `
}

// Error 回调错误信息
func Error(ctx *gin.Context, code int) {

	Data := &Data{
		Code:    code,
		Message: getMessage(code),
	}

	ctx.JSON(http.StatusOK, Data)
}

// Success 回调正确信息
func Success(ctx *gin.Context, data interface{}) {
	Data := &Data{
		Code:    ErrCodeSuccess,
		Message: getMessage(ErrCodeSuccess),
		Data:    data,
	}

	ctx.JSON(http.StatusOK, Data)
}
`
