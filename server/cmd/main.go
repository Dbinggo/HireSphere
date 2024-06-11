package main

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/internal/db/myRedis"
	"github.com/Dbinggo/HireSphere/server/internal/db/mySql"
	"github.com/Dbinggo/HireSphere/server/log"
)

func main() {
	log.InitLogger()
	configs.InitConfig()
	mySql.InitMySql()
	myRedis.InitMyRedis()

}
