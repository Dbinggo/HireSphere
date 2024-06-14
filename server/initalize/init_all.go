package initalize

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/utils"
)

func Init() {
	InitPath()
	InitConfig()
	InitLog()
	//InitDataBase()
}
func InitPath() {
	global.Path = utils.GetRootPath("")
}
