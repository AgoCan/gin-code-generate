package generators

import (
	"fmt"
	"os"
	"text/template"
)

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

// GeneratorMgr 生成器管理
type GeneratorMgr struct {
	genMap map[string]Generator
}

var genMgr *GeneratorMgr

// Generator 生成器接口
type Generator interface {
	// Run 生成器run
	Run(opt *Option) error
}

// Register 把生成器都注册到map中，然后轮询执行
func Register(name string, gen Generator) (err error) {
	_, ok := genMgr.genMap[name]
	if ok {
		err = fmt.Errorf("genrator %v exits", name)
	}
	genMgr.genMap[name] = gen
	return nil
}

// writeFile 使用模版文件直接写入文件
func writeFile(tmplContent, filePath string, opt *Option) (err error) {

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// t, err := template.ParseFiles(tmplFilePath)
	t, err := template.New(filePath).Parse(tmplContent)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	err = t.Execute(file, &opt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// init 初始化
func init() {
	genMgr = &GeneratorMgr{
		genMap: make(map[string]Generator),
	}
}

// RunGenerator 运行所有已经注册的生成器
func RunGenerator(opt *Option) (err error) {
	for _, gen := range genMgr.genMap {
		err = gen.Run(opt)
		return err
	}
	return nil
}
