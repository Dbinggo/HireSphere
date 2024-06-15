package initalize

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/db/databases"
	"github.com/Dbinggo/HireSphere/server/db/myRedis"
	"github.com/Dbinggo/HireSphere/server/global"
)

func InitDataBase() {
	conf := configs.Conf.DB
	switch conf.Driver {
	case "mysql":
		mysql := &databases.Mysql{}
		db, err := mysql.InitDataBases()
		if err != nil {
			global.Logger.Panic("mysql数据库初始化失败！")
		}
		global.DB = db
		break
	}
	if global.Config.App.Env != "pro" {
		err := global.DB.AutoMigrate()
		if err != nil {
			global.Logger.Panic("数据库迁移失败！")
		}
	}
	global.Logger.Info("数据库初始化成功！")
}
func InitRedis() {
	if global.Config.Redis.Enable {
		var err error
		global.Rdb, err = myRedis.GetRedisClient()
		if err != nil {
			global.Logger.Errorf("无法初始化Redis : %v", err)
		}
	} else {
		global.Logger.Warn("不使用Redis")
	}

}
