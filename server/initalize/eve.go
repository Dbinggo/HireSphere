package initalize

import "github.com/Dbinggo/HireSphere/server/global"

func Eve() {
	global.Rdb.Close()
	sqlDB, _ := global.DB.DB()
	sqlDB.Close()
}
