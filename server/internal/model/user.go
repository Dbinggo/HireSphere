package model

import (
	"fmt"
	"github.com/Dbinggo/HireSphere/server/global"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	ID       int64  `json:"id" gorm:"primary_key" redis:"id" gorm:"column:id;" `
	Username string `json:"username" redis:"username" gorm:"column:username;" `
	Password string `json:"password" redis:"password" gorm:"column:password;" `
	Email    string `json:"email" redis:"email" gorm:"column:email;"`
	Phone    string `json:"phone" redis:"phone" gorm:"column:phone;"`
	Role     int64  `json:"role" redis:"role" gorm:"column:role;"`
	Status   int64  `json:"status" redis:"status" gorm:"column:status;"`
}

// TableName 表名
func (User) TableName() string {
	return "user"
}

// Where 查询条件
func (u User) Where() *gorm.DB {
	return global.DB.Model(&u).Where("id = ?", u.ID)
}

// KeyName RedisKey
func (u User) KeyName() string {
	return fmt.Sprintf("user:%d:hash", u.ID)
}
