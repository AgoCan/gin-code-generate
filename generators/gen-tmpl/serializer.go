package tmpl

// SerializerContent 回调使用
var SerializerContent = `package serializer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 回调时的固定内容
type Response struct {
	Code    int         ` + "`json:\"code\"`" + `
	Message string      ` + "`json:\"message\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
}

// Error 回调错误信息
func Error(code int) (response Response) {
	resData := &resData{
		Code:    code,
		Message: getMessage(code),
	}
}

// Success 回调正确信息
func Success(data interface{}) (response Response) {
	resData := &resData{
		Code:    CodeSuccess,
		Message: getMessage(CodeSuccess),
		Data:    data,
	}
}
`
