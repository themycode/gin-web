package user_request

type FindUserById struct {
	ID int `json:"id" binding:"required"`
}