package generators

import (
	"fmt"
	"path"

	tmpl "github.com/agocan/gin-code-generate/generators/gen-tmpl"
)

var files = map[string]string{
	"main.go":                        tmpl.MainContent,
	"routers/router.go":              tmpl.RouterContent,
	"middleware/log.go":              tmpl.MiddlewareLog,
	"config/config.go":               tmpl.ConfigContent,
	"config/config.yaml":             tmpl.ConfigYamlContent,
	"models/model.go":                tmpl.ModelContent,
	"Dockerfile":                     tmpl.DockerfileContent,
	"README.md":                      tmpl.ReadmeContent,
	"config/option.go":               tmpl.ConfigOptionContent,
	"utils/generator/id.go":          tmpl.IDGenerateContent,
	"utils/response/response.go":     tmpl.ResponseContent,
	"utils/response/responsecode.go": tmpl.ResponseCodeContent,
	".gitignore":                     tmpl.Ignore,
}

// FileGenerator 文件生成器
type FileGenerator struct {
}

// FileGen 文件生成器实例
var FileGen *FileGenerator

// Run 运行生成器
func (f *FileGenerator) Run(opt *Option) error {
	for fileName, tmplContent := range files {
		filePath := path.Join(opt.AbsProjectPath, fileName)

		err := writeFile(tmplContent, filePath, opt)
		if err != nil {
			fmt.Printf("create files err：%v", err)
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

	err := writeFile(tmpl.GoModContent, filePath, opt)
	if err != nil {
		return err
	}
	return nil
}
