package test

import (
	"context"
	"github.com/Dbinggo/HireSphere/server/common/databases"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/initalize"
	"github.com/Dbinggo/HireSphere/server/internal/model"
	"testing"
)

func TestFindInRedisOrInDB(t *testing.T) {
	initalize.Init()
	user := model.User{
		ID:       1,
		Username: "小明",
		Password: "123",
		Email:    "123",
		Phone:    "123",
		Role:     1,
		Status:   1,
	}
	global.DB.Create(user)

	var findUserRdb = func() error {
		return global.Rdb.HGetAll(context.Background(), user.KeyName()).Scan(&user)
	}
	var findUserDB = func() error {
		return global.DB.Where("id = ?", user.ID).First(&user).Error
	}
	var setUserRdb = func() error {
		return global.Rdb.HMSet(context.Background(), user.KeyName(), user).Err()
	}

	err := databases.FindInRedisOrDB(context.Background(), findUserRdb, findUserDB, setUserRdb, user)
	if err != nil {
		t.Errorf("FindInRedisOrDB() error = %v", err)
		return
	}

	t.Log(user)
}
