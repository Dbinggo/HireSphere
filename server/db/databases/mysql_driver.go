package databases

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
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
		global.Logger.Panic("无法连接数据库！: %v", err)
		return nil, err
	}
	global.Logger.Info("数据库连接成功！")
	return db, nil
}
func (m *Mysql) getDsn() string {
	return configs.Conf.DB.Dsn
}
