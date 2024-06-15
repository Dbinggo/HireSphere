package initalize

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/utils"
)

func Init() {
	InitLog()
	InitPath()
	InitConfig()
	InitLog()
	InitDataBase()
	InitRedis()
}
func InitPath() {
	global.Path = utils.GetRootPath("")
}
