package global

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Path   string
	DB     *gorm.DB
	Rdb    *redis.Client
	Config *configs.Config
)
