package main

import (
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/initalize"
	"github.com/Dbinggo/HireSphere/server/internal/router"
)

func main() {

	initalize.Init()
	// 工程进入前夕，释放资源
	defer initalize.Eve()
	router.RunServer(*global.Config)
	zlog.Infof("program is end！")

}
