package generators

// 创建目录需要单独运行，不放在GeneratorMgr里面
import (
	"os"
	"path"
)

var dirs = []string{
	"handlers",
	"routers",
	"models",
	"templates",
	"utils/response",
	"utils/generator",
	"config",
	"log",
	"middleware",
}

// DirGenerator 目录生成器
type DirGenerator struct {
}

// Run 运行生成器
func (d *DirGenerator) Run(opt *Option) (err error) {
	for _, dir := range dirs {
		fullDir := path.Join(opt.AbsProjectPath, dir)
		err = os.MkdirAll(fullDir, 0755)
		if err != nil {
			panic(err)

		}
	}
	return nil
}
