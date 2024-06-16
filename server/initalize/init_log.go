package initalize

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/log"
	"github.com/Dbinggo/HireSphere/server/log/zlog"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitLog(config *configs.Config) {
	zapLogger, hertzLogger := log.GetLogger(config)
	zlog.InitLogger(zapLogger)
	hlog.SetLogger(hertzLogger)
}
