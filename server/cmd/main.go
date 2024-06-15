package main

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/initalize"
)

func main() {

	initalize.Init()

	global.Logger.Info("程序运行成功！")
}
