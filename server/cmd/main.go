package main

import (
	"github.com/Dbinggo/HireSphere/server/initalize"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {

	initalize.Init()
	opt := server.WithHostPorts(":8080")
	h := server.Default(opt)
	h.Spin()
}
