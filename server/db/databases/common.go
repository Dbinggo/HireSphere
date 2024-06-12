package databases

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/sirupsen/logrus"
)

type DataBase interface {
	getDsn() string
	InitDataBases() error
}

func InitDataBase() {
	conf := configs.Conf.DB
	switch conf.Driver {
	case "mysql":
		mysql := &Mysql{}
		err := mysql.InitDataBases()
		if err != nil {
			logrus.Fatal("mysql数据库初始化失败！")
		}
		break
	}

}
