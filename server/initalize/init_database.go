package initalize

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/db/databases"
	"github.com/Dbinggo/HireSphere/server/db/myRedis"
	"github.com/Dbinggo/HireSphere/server/global"
)

func InitDataBase(config configs.Config) {
	switch config.DB.Driver {
	case "mysql":
		databases.InitDataBases(databases.NewMySql(), config)
		break
	}
	if config.App.Env != "pro" {
		err := global.DB.AutoMigrate()
		if err != nil {
			global.Logger.Panic("数据库迁移失败！")
		}
	}
	global.Logger.Info("数据库初始化成功！")
}
func InitRedis(config configs.Config) {
	if config.Redis.Enable {
		var err error
		global.Rdb, err = myRedis.GetRedisClient(config)
		if err != nil {
			global.Logger.Errorf("无法初始化Redis : %v", err)
		}
	} else {
		global.Logger.Warn("不使用Redis")
	}

}
