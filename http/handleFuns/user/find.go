package user_http

import (
	"github.com/CoderSamYhc/gin-web/entity"
	"github.com/CoderSamYhc/gin-web/http"
	"github.com/CoderSamYhc/gin-web/http/requests"
	"github.com/CoderSamYhc/gin-web/models"
	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context)  {
	var (
		err error
		params entity.FindUser
	)

	err = c.ShouldBindJSON(&params)
	if err != nil {
		http.ErrorResponse(c, 10000, requests.Translate(err))
		return
	}

	user := models.User{}
	result, err := user.FindById(params.ID)
	if err != nil {
		http.ErrorResponse(c, 10000, err)
		return
	}
	http.SuccessResponse(c, result)
}
