package main

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/db/databases"
	"github.com/Dbinggo/HireSphere/server/db/myRedis"
	"github.com/Dbinggo/HireSphere/server/log"
)

func main() {
	log.InitLogger()
	configs.InitConfig()
	databases.InitDataBase()
	myRedis.InitMyRedis()

}
