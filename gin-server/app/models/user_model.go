//@Author: wulinlin
//@Description:
//@File:  user_model
//@Version: 1.0.0
//@Date: 2024/01/02 22:00

package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	// 可以添加其他用户信息字段
}
