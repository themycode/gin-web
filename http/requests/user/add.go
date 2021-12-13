package user_request

type AddUser struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age int8 `form:"age" json:"age" binding:"required"`
	Sex int8 `form:"sex" json:"sex" binding:"required,oneof=1 2"`
}