package user_http

import (
	"github.com/CoderSamYhc/gin-web/entity"
	"github.com/CoderSamYhc/gin-web/http"
	"github.com/CoderSamYhc/gin-web/http/requests"
	"github.com/CoderSamYhc/gin-web/models"
	"github.com/gin-gonic/gin"
)



func Add(c *gin.Context)  {
	var (
		params entity.User
		err error
	)
	err = c.ShouldBindJSON(&params)
	if err != nil {
		http.ErrorResponse(c, 10000, requests.Translate(err))
		return
	}

	// 调model入库
	user := models.User{}
	id, err := user.Add(&params)
	if err != nil {
		http.ErrorResponse(c, 10000, err.Error())
		return
	}

	http.SuccessResponse(c, id)
	return
}