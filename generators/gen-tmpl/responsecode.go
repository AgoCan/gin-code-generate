package tmpl

// ResponseCodeContent 回调的错误码
var ResponseCodeContent = `package response

// 错误码
const (
	ErrCodeSuccess           = 0
	ErrCodeParameter         = 1001
)

func getMessage(code int) (message string) {
	switch code {
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeParameter:
		message = "参数错误"
	default:
		message = "记录已经存在"
		message = "未知错误"
	}
	return
}
`
