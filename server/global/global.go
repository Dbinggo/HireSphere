package global

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	//跟目录下
	Path   string
	DB     *gorm.DB
	Rdb    *redis.Client
	Config *configs.Config
	Log    *zap.SugaredLogger
)
