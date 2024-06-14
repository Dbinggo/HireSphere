package main

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/initalize"
)

func main() {

	initalize.Init()
	//opt := server.WithHostPorts(":8080")
	//h := server.Default(opt)
	//h.Spin()
	global.Logger.Info("程序运行成功！")
}
