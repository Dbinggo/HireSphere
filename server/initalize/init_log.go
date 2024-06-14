package initalize

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/log"
)

func InitLog() {
	logger := log.GetZap()
	global.Logger = logger.Sugar()
	global.Logger.Info("path is ", global.Path)
}
