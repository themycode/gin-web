package models

import (
	"errors"
	"github.com/CoderSamYhc/gin-web/db"
	user_request "github.com/CoderSamYhc/gin-web/http/requests/user_request"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
	Sex int8 `json:"sex"`
}

func (u *User) TableName() string {
	return "user_request"
}

func (u *User) Add(params *user_request.AddUser) (int, error) {
	data := User{
		Name: params.Name,
		Age: params.Age,
		Sex: params.Sex,
	}
	result := db.GetDB().Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return data.ID, nil
}

func (u *User) FindById(id int) (*User, error) {
	user := &User{}
	result := db.GetDB().Where(id).Find(user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("数据不存在")
	}

	return user, nil
}