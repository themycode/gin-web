package models

import (
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
