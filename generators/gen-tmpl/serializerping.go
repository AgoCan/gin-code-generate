package tmpl

// SerializerPingContent 回调的错误码
var SerializerPingContent = `package serializer

import "{{ .ProjectName }}/models"

// Ping 测试序列化器
type Ping struct {
	ID  int    ` + "`json:\"id\"`" + `
	Msg string ` + "`json:\"msg\"`" + `
}

//BuildPing 测试序列化器
func BuildPing(ping models.Ping) Ping {
	return Ping{
		ID:  ping.ID,
		Msg: ping.Msg,
	}
}

`
