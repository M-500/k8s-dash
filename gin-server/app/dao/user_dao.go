//@Author: wulinlin
//@Description:
//@File:  user_dao
//@Version: 1.0.0
//@Date: 2024/01/02 22:00

package dao

import "gin-server/app/models"

type UserDao struct {
	db *xorm.DB
}

func NewUserDao(db *xorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) CreateUser(user *models.User) error {
	return dao.db.Create(user).Error
}

func (dao *UserDao) GetUserByUsername(username string) (*models.User, error) {
	user := &model.User{}
	err := dao.db.Where("username = ?", username).First(user).Error
	return user, err
}
