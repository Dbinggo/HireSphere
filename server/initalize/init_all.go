package initalize

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/utils"
)

func Init() {
	InitLog(global.Config)
	InitPath()
	InitConfig()
	InitLog(global.Config)
	InitDataBase(*global.Config)
	InitRedis(*global.Config)
}
func InitPath() {
	global.Path = utils.GetRootPath("")
}
