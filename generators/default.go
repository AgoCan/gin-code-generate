package generators

import (
	"fmt"
	"io/ioutil"
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
	config/config_test.go
	config/configstruct.go
	middleware/log.go
	routers/router.go
	model/model.go
*/


// DefaultGenerator 默认生成器
func DefaultGenerator(opt *Option) (err error) {
	// prePath 指定路径生成项目
	var dirs []string
	dirs = []string{"handlers", "routers", "model", "templates",
		"utils/logging", "config", "log", "middleware"}
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

// writeFile 写入文件，因为比较小。直接一次性写入
func writeFile(filePath, content string)(){
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		// 如果错误。直接抛错
		fmt.Println("write file failed, err:", err)
		panic(err)
	}
}

// DefaultFileGenerator 创建文件
func DefaultFileGenerator(opt *Option)(err error){
	// 入口函数
	mainFile := path.Join(opt.AbsProjectPath, "main.go")
	mainContent := fmt.Sprintf(MainContent, opt.ProjectName, opt.ProjectName)
	writeFile(mainFile, mainContent)

	// 写路由
	routerFile := path.Join(opt.AbsProjectPath, "routers/router.go")
	routerContent := fmt.Sprintf(RouterContent, opt.ProjectName)
	writeFile(routerFile, routerContent)

	// 日志中间件
	middlewareLogFile  := path.Join(opt.AbsProjectPath, "middleware/log.go")
	middlewareLogContent := fmt.Sprintf(MiddlewareLogContent,opt.ProjectName, opt.ProjectName)
	writeFile(middlewareLogFile, middlewareLogContent)

	utilsLogFile := path.Join(opt.AbsProjectPath, "utils/logging/log.go")
	utilsLogContent := fmt.Sprintf(UtilsLogContent, opt.ProjectName)
	writeFile(utilsLogFile, utilsLogContent)

	// 配置相关
	configFile := path.Join(opt.AbsProjectPath, "config/config.go")
	writeFile(configFile, ConfigContent)
	configYamlFile := path.Join(opt.AbsProjectPath, "config/config.yaml")
	writeFile(configYamlFile, ConfigYamlContent)
	configTestFile := path.Join(opt.AbsProjectPath, "config/config_test.go")
	writeFile(configTestFile, ConfigTestContent)
	configStructFile := path.Join(opt.AbsProjectPath, "config/configstruct.go")
	writeFile(configStructFile, ConfigStructContent)

	modelFile := path.Join(opt.AbsProjectPath, "model/model.go")
	modelContent := fmt.Sprintf(ModelContent, opt.ProjectName)
	writeFile(modelFile, modelContent)
	return nil
}

func DefaultModGenerator(opt *Option)(err error){
	goModFile := path.Join(opt.AbsProjectPath, "go.mod")
	goModContent := fmt.Sprintf(GoModContent, opt.ProjectName)
	writeFile(goModFile, goModContent)
	return nil
}