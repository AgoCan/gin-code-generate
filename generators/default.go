package generators

import (
	"fmt"
	"os"
	"path"
)

/*
创建基础目录
	handlers
	routers
	models
	templates
	utils
	config
	log
	middleware
创建基础文件
	main.go
	Dockerfile
	README.md
	config/config.go
	config/config.yaml
	config/config.yaml
	config/config_test.go
	config/configstruct.go
	middleware/log.go
	routers/router.go
*/


// DefaultGenerator 默认生成器
func DefaultGenerator(opt *Option) (err error) {
	// prePath 指定路径生成项目
	fmt.Println(opt)
	var dirs []string
	dirs = []string{"handlers", "routers", "models", "templates",
		"utils", "config", "log", "middleware"}
	for _, dir := range dirs {
		fullDir := path.Join(opt.AbsProjectPath, dir)
		err = os.MkdirAll(fullDir, 0755)
		if err != nil {
			panic(err)
			return err
		}
	}
	return nil
}
