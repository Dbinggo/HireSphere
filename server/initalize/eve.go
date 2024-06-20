package initalize

import (
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/global"
	"runtime"
)

func Eve() {
	zlog.Warnf("start to end！")
	errRedis := global.Rdb.Close()
	if errRedis != nil {
		zlog.Errorf("failed to close redis with err ：%v", errRedis.Error())
	}

	sqlDB, _ := global.DB.DB()
	errDB := sqlDB.Close()
	if errDB != nil {
		zlog.Errorf("failed to close db with err:%v", errDB.Error())
	}
	runtime.GC()
	if errDB == nil && errRedis == nil {
		zlog.Warnf("end success!")
	}
}
