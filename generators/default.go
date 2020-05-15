package generators

import (
	"os"
	"strings"
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
func DefaultGenerator(path string, projectPath string) {
	var dirs []string
	dirs = []string{"handlers", "routers", "models", "templates",
		"utils", "config", "log", "middleware"}
	for _, dir := range dirs {
		fullDir := strings.Trim(path, string(os.PathSeparator)) + string(os.PathSeparator) +
			projectPath + string(os.PathSeparator) + dir
		err := os.MkdirAll(fullDir, 0755)
		if err != nil {
			panic(err)
			return
		}
	}
}
