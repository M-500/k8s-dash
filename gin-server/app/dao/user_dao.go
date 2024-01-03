//@Author: wulinlin
//@Description:
//@File:  user_dao
//@Version: 1.0.0
//@Date: 2024/01/02 22:00

package dao

import (
	"context"
	"errors"
	"gin-server/app/models"
	"gin-server/pkg/utils/db_helper"
	"time"
	"xorm.io/xorm"
)

type UserDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		db:  db_helper.GetDb(),
		ctx: ctx,
	}
}

func (dao *UserDao) Get(id int) (*models.User, error) {
	queryData := &models.User{}
	exist, err := dao.db.ID(id).Get(queryData)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("数据不存在")
	}
	return queryData, nil
}
func (dao *UserDao) FindByUserName(username string) (*models.User, error) {
	queryData := &models.User{}
	exist, err := dao.db.Where("`username` = ?", username).Get(queryData)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("数据不存在")
	}
	return queryData, nil
}
func (dao *UserDao) FindByPhone(phone string) (*models.User, error) {
	data := &models.User{}
	res, err := dao.db.Where("`email` = ?", phone).Get(data)
	if err != nil {
		return nil, err
	}
	if !res {
		return nil, errors.New("数据不存在")
	}
	return data, nil
}
func (dao *UserDao) FindAllPager(username string, page int, size int) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	start := (page - 1) * size
	dataList := make([]models.User, 0)
	total, err := dao.db.Where("`username` like ?", username).Desc("id").Limit(size, start).FindAndCount(&dataList)
	return dataList, total, err
}
func (dao *UserDao) Insert(data *models.User) error {
	now := time.Now()
	data.CreateTime = &now
	data.UpdateTime = &now
	_, err := dao.db.Insert(data)
	return err
}

func (dao *UserDao) Update(data *models.User, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	now := time.Now()
	data.UpdateTime = &now
	_, err := sess.Update(data)
	return err
}

// 更高一层的封装
func (dao *UserDao) Save(data *models.User, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	}
	return dao.Insert(data)
}
