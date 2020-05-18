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

var ErrIsNotDir = errors.New("the path is not director")
var opt generators.Option

// 入口函数
func entry(c *cli.Context) (err error) {
	// 判断路径是否存在

	s, err := os.Stat(c.String("path"))

	if err != nil {
		fmt.Printf("'path' %v is no such file or director.", s)
		return err
	}
	if !s.IsDir(){
		panic("the path is not director")
		return ErrIsNotDir
	}
	opt.AbsProjectPath = path.Join(c.String("path"), c.String("project-name"))

	err = generators.DefaultGenerator(&opt)
	if err != nil {
		fmt.Printf("create dirs err: %v", err)
	}

	return nil
}

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Value:   ".",
				Usage:   "生成项目的路径",
				Destination: &opt.ProjectPath,
			},
			&cli.StringFlag{
				Name:    "project-name",
				Value:   "demo",
				Usage:   "项目名称",
				Destination: &opt.ProjectName,
			},

		},

		Action: entry,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
