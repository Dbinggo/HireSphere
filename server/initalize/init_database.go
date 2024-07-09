package initalize

import (
	"github.com/Dbinggo/HireSphere/server/common/databases"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/common/myRedis"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/internal/model"
)

func InitDataBase(config configs.Config) {
	switch config.DB.Driver {
	case "mysql":
		databases.InitDataBases(databases.NewMySql(), config)
		break
	}
	if config.App.Env != "pro" {
		err := global.DB.AutoMigrate(&model.User{})
		if err != nil {
			zlog.Fatalf("数据库迁移失败！")
		}
	}
	zlog.Infof("数据库初始化成功！")
}
func InitRedis(config configs.Config) {
	if config.Redis.Enable {
		var err error
		global.Rdb, err = myRedis.GetRedisClient(config)
		if err != nil {
			zlog.Errorf("无法初始化Redis : %v", err)
		}
	} else {
		zlog.Warnf("不使用Redis")
	}

}
