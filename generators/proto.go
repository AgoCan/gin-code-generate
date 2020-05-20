package generators

import (
	"fmt"
	"os"
	"os/exec"
)

func ProtoGenerator(opt *Option) (err error) {
	cmd := exec.Command("protoc", "--go_out=plugins=grpc:.", opt.ProtoFilePath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("proto err: %s" , err)
		return
	}
	return nil
}