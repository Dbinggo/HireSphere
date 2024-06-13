package global

import (
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Redis  *redis.Client
	Config *configs.Config
)
