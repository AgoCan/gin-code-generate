package generators

// Option 参数保存
type Option struct {
	// AbsProjectPath 项目路径+项目名称
	AbsProjectPath string
	// ProjectPath 项目路径
	ProjectPath string
	// ProjectName 项目名称
	ProjectName string
	// proto 路径
	ProtoFilePath string
	//
	IsMod bool
}

// Generator 生成器接口
type Generator interface {
	Run(opt *Option) error
}
