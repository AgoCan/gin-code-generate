package generators

import "os/exec"

func ProtoGenerator(projectPath,filename string) (err error) {
	exec.Command("protoc", "--go_out", )
	return nil
}