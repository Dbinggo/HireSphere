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

	err := databases.FindInRedisOrDB(context.Background(), &user)
	if err != nil {
		t.Errorf("FindInRedisOrDB() error = %v", err)
		return
	}

	t.Log(user)
}
