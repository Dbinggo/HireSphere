package main

import (
	"github.com/Dbinggo/HireSphere/server/initalize"
	"github.com/Dbinggo/HireSphere/server/internal/router"
	"github.com/Dbinggo/HireSphere/server/log/zlog"
)

func main() {

	initalize.Init()
	// 工程进入前夕，释放资源
	defer initalize.Eve()
	router.RunServer()
	zlog.Infof("程序运行成功！")

}
