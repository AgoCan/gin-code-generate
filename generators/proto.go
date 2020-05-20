package generators

import (
	"fmt"
	"os/exec"
)

func ProtoGenerator(opt *Option) (err error) {
	cmd := exec.Command("protoc", "--go_out=plugins=grpc:.", opt.ProtoFilePath)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("proto err: %s" , err)
		return
	}
	return nil
}