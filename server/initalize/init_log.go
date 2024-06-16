package initalize

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/log"
	"github.com/Dbinggo/HireSphere/server/log/zlog"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitLog(config *configs.Config) {
	logger := log.GetZap(config)
	zlog.InitLogger(logger)
	hlog.SetLogger(logger)
}
