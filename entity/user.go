package entity

type User struct {
	ID int `json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
	Age int8 `form:"age" json:"age" binding:"required"`
	Sex int8 `form:"sex" json:"sex" binding:"required,oneof=1 2"`
}

type FindUser struct {
	ID int `json:"id" binding:"required"`
}