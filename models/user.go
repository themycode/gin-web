package models

import (
	"errors"
	"github.com/CoderSamYhc/gin-web/db"
	"github.com/CoderSamYhc/gin-web/entity"
)

type User struct {}

func (u *User) Add(data *entity.User) (int, error) {
	result := db.GetDB().Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return data.ID, nil
}

func (u *User) FindById(id int) (*entity.User, error) {
	user := &entity.User{}
	result := db.GetDB().Where(id).Find(user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("数据不存在")
	}

	return user, nil
}