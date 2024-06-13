package databases

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
}

// InitDataBases 初始化
func (m *Mysql) InitDataBases() (*gorm.DB, error) {
	dsn := m.getDsn()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.Fatalf("无法连接数据库！: %v", err)
		return nil, err
	}
	logrus.Infof("数据库连接成功！")
	return db, nil
}
func (m *Mysql) getDsn() string {
	return configs.Conf.DB.Dsn
}
