package tmpl

// ModelPingContent model
var ModelPingContent = `package models

// Ping 测试
type Ping struct {
	baseModel
	Msg string ` + "`db:\"msg\" json:\"msg\"`" + `
}

`
