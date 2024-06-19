package initalize

import (
	"github.com/Dbinggo/HireSphere/server/common/log"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitLog(config *configs.Config) {
	zapLogger, hertzLogger := log.GetLogger(config)
	zlog.InitLogger(zapLogger)
	hlog.SetLogger(hertzLogger)
	hlog.Debugf("for test for test for test")
	hlog.Infof("for test for test for test")
	hlog.Warnf("for test for test for test")
	hlog.Errorf("for test for test for test")
}
