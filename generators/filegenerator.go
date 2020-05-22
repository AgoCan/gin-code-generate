package generators

import (
	"path"
)

var files = []string{
	"main.go",
	"routers/router.go",
	"middleware/log.go",
	"utils/logging/log.go",
	"config/config.go",
	"config/config.yaml",
	"config/config_test.go",
	"config/configstruct.go",
	"model/model.go",
	"Dockerfile",
	"README.md",
}

// FileGenerator 文件生成器
type FileGenerator struct {
}

// FileGen 文件生成器实例
var FileGen *FileGenerator

// Run 运行生成器
func (f *FileGenerator) Run(opt *Option) error {
	for _, fileName := range files {
		filePath := path.Join(opt.AbsProjectPath, fileName)
		tmplFilePath := path.Join("./templates/", fileName+".tmpl")

		err := writeFile(tmplFilePath, filePath, opt)
		if err != nil {
			return err
		}
	}

	return nil
}

// ModGenerator 文件生成器
type ModGenerator struct {
}

// ModGen 文件生成器实例
var ModGen *ModGenerator

// Run 运行生成器
func (m *ModGenerator) Run(opt *Option) error {
	filePath := path.Join(opt.AbsProjectPath, "go.mod")
	tmplFilePath := path.Join("./templates/", "go.mod.tmpl")
	err := writeFile(tmplFilePath, filePath, opt)
	if err != nil {
		return err
	}
	return nil
}
