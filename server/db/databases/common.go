package databases

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"gorm.io/gorm"
)

type DataBase interface {
	getDsn(config configs.Config) string
	initDataBases(config configs.Config) (*gorm.DB, error)
}

func InitDataBases(base DataBase, config configs.Config) {
	var err error
	global.DB, err = base.initDataBases(config)
	if err != nil {
		global.Logger.Fatal("无法初始化数据库 %v", err)
		return
	}
	global.Logger.Info("初始化数据库成功！")
	return
}
