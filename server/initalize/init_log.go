package initalize

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/log"
)

func InitLog(config *configs.Config) {
	logger := log.GetZap(config)
	global.Logger = logger.Sugar()
}
