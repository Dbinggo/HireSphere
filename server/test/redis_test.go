package test

import (
	"context"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/initalize"
	"testing"
)

func TestRedis(T *testing.T) {
	initalize.Init()
	global.Rdb.Set(context.Background(), "test_redis", "1111", 0)
	v1 := global.Rdb.Get(context.Background(), "test_redis")
	global.Rdb.Set(context.Background(), "test_redis", "2222", 0)
	v2 := global.Rdb.Get(context.Background(), "test_redis")
	global.Rdb.Set(context.Background(), "test_redis", "3333", 0)
	v3 := global.Rdb.Get(context.Background(), "test_redis")
	global.Logger.Info(v1, v2, v3)
}
