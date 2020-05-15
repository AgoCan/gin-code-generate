package main

import (
	"errors"
	"log"
	"os"

	"github.com/agocan/gin-code-generate/generators"
	"github.com/urfave/cli/v2"
)

var ErrIsNotDir = errors.New("the path is not director")

// 入口函数
func entry(c *cli.Context) (err error) {
	// 判断路径是否存在
	s,err := os.Stat(c.String("path"))
	if err != nil {
		panic("'path' is no such file or director.")
		return err
	}
	if !s.IsDir(){
		panic("the path is not director")
		return ErrIsNotDir
	}
	generators.DefaultGenerator(c.String("path"), c.String("project-path"))

	return nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Value:   ".",
				Usage:   "生成项目的路径",
			},
			&cli.StringFlag{
				Name:    "project-path",
				Value:   "demo",
				Usage:   "项目名称",
			},

		},

		Action: entry,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
