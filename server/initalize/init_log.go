package initalize

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/log"
)

func InitLog() {
	logger := log.GetZap()
	global.Log = logger.Sugar()
	global.Log.Info("path is ", global.Path)
}
