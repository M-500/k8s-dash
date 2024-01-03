//@Author: wulinlin
//@Description:
//@File:  user_service
//@Version: 1.0.0
//@Date: 2024/01/02 22:01

package services

import (
	"context"
	"gin-server/app/dao"
	"gin-server/app/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDao *dao.UserDao
	ctx     context.Context
}

func NewUserService(userDao *dao.UserDao, ctx context.Context) *UserService {
	return &UserService{
		userDao: userDao,
		ctx:     ctx,
	}
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

	return service.userDao.Insert(user)
}

func (service *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := service.userDao.FindByUserName(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
