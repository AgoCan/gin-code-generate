package tmpl

// ResponseContent 回调使用
var ResponseContent = `package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// resData 回调时的固定内容
type resData struct {
	Code    int         ` + "`json:" + `"code"` + "`" + `
	Message string      ` + "`json:" + `"message"` + "`" + `
	Data    interface{} ` + "`json:" + `"data"` + "`" + `
}

// Error 回调错误信息
func Error(ctx *gin.Context, code int) {

	resData := &resData{
		Code:    code,
		Message: getMessage(code),
	}

	ctx.JSON(http.StatusOK, resData)
}

// Success 回调正确信息
func Success(ctx *gin.Context, data interface{}) {
	resData := &resData{
		Code:    ErrCodeSuccess,
		Message: getMessage(ErrCodeSuccess),
		Data:    data,
	}

	ctx.JSON(http.StatusOK, resData)
}
`
