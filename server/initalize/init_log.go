package initalize

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/log"
	"github.com/Dbinggo/HireSphere/server/log/zlog"
)

func InitLog(config *configs.Config) {
	logger := log.GetZap(config)
	zlog.InitLogger(logger)
}
