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
func (m *Mysql) initDataBases(config configs.Config) (*gorm.DB, error) {
	dsn := m.getDsn(config)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		global.Logger.Panic("MySQL无法连接数据库！: %v", err)
		return nil, err
	}
	global.Logger.Info("MySQL连接数据库成功！")
	return db, nil
}
func (m *Mysql) getDsn(config configs.Config) string {
	return config.DB.Dsn
}
func NewMySql() DataBase {
	return &Mysql{}
}
