package initalize

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/log"
)

func InitLog() {
	logger := log.GetZap()
	global.Log = logger.Sugar()

	global.Log.Infof("hello")
	global.Log.Errorf("hello")

}
