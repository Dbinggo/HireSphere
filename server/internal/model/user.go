package model

// User 用户表
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
	Status   int    `json:"status"`
}

// TableName 表名
func (User) TableName() string {
	return "user"
}
