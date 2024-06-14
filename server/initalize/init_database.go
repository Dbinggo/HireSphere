package initalize

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/db/databases"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/sirupsen/logrus"
)

func InitDataBase() {
	conf := configs.Conf.DB
	switch conf.Driver {
	case "mysql":
		mysql := &databases.Mysql{}
		db, err := mysql.InitDataBases()
		if err != nil {
			global.Log.Panic("mysql数据库初始化失败！")
		}
		global.DB = db
		break
	}
	if global.Config.App.Env != "pro" {
		err := global.DB.AutoMigrate()
		if err != nil {
			logrus.Fatal("数据库迁移失败！")
		}
	}
	global.Log.Info("数据库初始化成功！")
}
