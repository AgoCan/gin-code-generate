package tmpl

// IDGenerateContent 传参的结构体
var IDGenerateContent = `package generator
// 唯一ID生成器
// 如果是分布式，根据settings里面配置好machineID即可
import "github.com/sony/sonyflake"

// IdGenerate 唯一ID生成器
func IdGenerate()(uint64, error){
	settings := sonyflake.Settings{}
	sk := sonyflake.NewSonyflake(settings)

	id, err := sk.NextID()
	return id,err
}
`
