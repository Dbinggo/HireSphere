package databases

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/internal/Model"
	"github.com/Dbinggo/HireSphere/server/log"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
}

// InitMySql 初始化
func (m *Mysql) InitDataBases() error {
	dsn := m.GetDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log.MyLogger,
	})
	if err != nil {
		logrus.Fatalf("无法连接数据库！: %v", err)
		return err
	}
	err = db.AutoMigrate(Model.User{})
	if err != nil {
		logrus.Fatalf("无法迁移数据库！: %v", err)
		return err
	}

	logrus.Infof("数据库连接成功！")
	return nil
}
func (m *Mysql) GetDsn() string {
	return configs.Conf.Dsn
}
