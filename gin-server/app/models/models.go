package models

import (
	"time"
)

type User struct {
	Id         uint       `xorm:"not null pk autoincr comment('主键id自增长') UNSIGNED INT"`
	Username   string     `xorm:"not null comment('用户名') VARCHAR(32)"`
	Password   string     `xorm:"not null comment('密码') VARCHAR(128)"`
	Email      string     `xorm:"comment('邮箱') VARCHAR(255)"`
	Phone      string     `xorm:"comment('手机号码') CHAR(11)"`
	NickName   string     `xorm:"comment('昵称') VARCHAR(32)"`
	Gender     int        `xorm:"comment('性别') TINYINT"`
	Age        int        `xorm:"comment('年龄') INT"`
	AvatarUrl  string     `xorm:"comment('头像') VARCHAR(255)"`
	CreateTime *time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime *time.Time `xorm:"not null comment('更新时间') DATETIME"`
	DeleteTime *time.Time `xorm:"comment('删除时间') DATETIME"`
}
