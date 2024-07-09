package test

import (
	"context"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/initalize"
	"github.com/davecgh/go-spew/spew"
	"github.com/redis/go-redis/v9"
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
	zlog.Infof("%v,%v,%v", v1, v2, v3)
}

type Model struct {
	Str1    string   `redis:"str1"`
	Str2    string   `redis:"str2"`
	Bytes   []byte   `redis:"bytes"`
	Int     int      `redis:"int"`
	Bool    bool     `redis:"bool"`
	Ignored struct{} `redis:"-"`
}

func TestRedis1(t *testing.T) {
	initalize.Init()
	rdb := global.Rdb
	ctx := context.Background()
	// Set some fields.
	if _, err := global.Rdb.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, "key", "str1", "hello")
		rdb.HSet(ctx, "key", "str2", "world")
		rdb.HSet(ctx, "key", "int", 123)
		rdb.HSet(ctx, "key", "bool", 1)
		rdb.HSet(ctx, "key", "bytes", []byte("this is bytes !"))
		return nil
	}); err != nil {
		panic(err)
	}

	var model1, model2, model3 Model

	// Scan all fields into the model.
	if err := rdb.HGetAll(ctx, "key").Scan(&model1); err != nil {
		panic(err)
	}

	// Or scan a subset of the fields.
	if err := rdb.HMGet(ctx, "key", "str1", "int").Scan(&model2); err != nil {
		panic(err)
	}
	if err := rdb.HSet(ctx, "key1", model1).Err(); err != nil {
		panic(err)
	}
	if err := rdb.HGetAll(ctx, "key1").Scan(&model3); err != nil {
		panic(err)
	}
	spew.Dump(model1)
	// Output:
	// (main.Model) {
	// 	Str1: (string) (len=5) "hello",
	// 	Str2: (string) (len=5) "world",
	// 	Bytes: ([]uint8) (len=15 cap=16) {
	// 	 00000000  74 68 69 73 20 69 73 20  62 79 74 65 73 20 21     |this is bytes !|
	// 	},
	// 	Int: (int) 123,
	// 	Bool: (bool) true,
	// 	Ignored: (struct {}) {
	// 	}
	// }

	spew.Dump(model2)
	// Output:
	// (main.Model) {
	// 	Str1: (string) (len=5) "hello",
	// 	Str2: (string) "",
	// 	Bytes: ([]uint8) <nil>,
	// 	Int: (int) 123,
	// 	Bool: (bool) false,
	// 	Ignored: (struct {}) {
	// 	}
	// }
	spew.Dump(model3)
}
