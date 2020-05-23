package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/agocan/gin-code-generate/generators"
	"github.com/urfave/cli/v2"
)

var (
	// ErrIsNotDir 错误
	ErrIsNotDir = errors.New("the path is not director")
	opt         generators.Option
)

// 入口函数
func entry(c *cli.Context) (err error) {
	// 判断路径是否存在

	s, err := os.Stat(c.String("path"))

	if err != nil {
		fmt.Printf("'path' %v is no such file or director.", s)
		return err
	}
	if !s.IsDir() {
		fmt.Printf("'path' %v is not director.", s)
		return ErrIsNotDir
	}
	opt.AbsProjectPath = path.Join(c.String("path"), c.String("project-name"))
	// 生成目录 需要在一开始就直接运行
	var dirGen generators.DirGenerator
	err = dirGen.Run(&opt)
	if err != nil {
		fmt.Printf("create dirs err: %v", err)
	}
	// 注册
	generators.Register("files",generators.FileGen)
	if c.Bool("mod") {
		generators.Register("mod",generators.ModGen)
	}
	generators.RunGenerator(&opt)

	return nil
}

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Value:       ".",
				Usage:       "生成项目的路径",
				Destination: &opt.ProjectPath,
			},
			&cli.StringFlag{
				Name:        "project-name",
				Value:       "demo",
				Usage:       "项目名称",
				Destination: &opt.ProjectName,
			},
			&cli.BoolFlag{
				Name:        "mod",
				Value:       true,
				Usage:       "生成mod",
				Destination: &opt.IsMod,
			},
		},

		Action: entry,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
