package initalize

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/log"
)

func InitLog() {
	logger := log.GetZap()
	global.Log = logger.Sugar()

	global.Log.Debug("hello")
	global.Log.Infof("hello")
	global.Log.Warn("hello")
	global.Log.Errorf("hello")
	global.Log.DPanic("hello")
	global.Log.Panic("hello")
}
