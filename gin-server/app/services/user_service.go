//@Author: wulinlin
//@Description:
//@File:  user_service
//@Version: 1.0.0
//@Date: 2024/01/02 22:01

package services

import (
	"gin-server/app/dao"
	"gin-server/app/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

func (service *UserService) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return service.userDao.CreateUser(user)
}

func (service *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := service.userDao.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
