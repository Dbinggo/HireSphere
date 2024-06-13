package main

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/db/databases"
	"github.com/Dbinggo/HireSphere/server/db/myRedis"
	"github.com/Dbinggo/HireSphere/server/log"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {

	log.InitLogger()
	configs.InitConfig()
	databases.InitDataBase()
	myRedis.InitMyRedis()
	opt := server.WithHostPorts(":8080")
	h := server.Default(opt)
	h.Spin()
}
