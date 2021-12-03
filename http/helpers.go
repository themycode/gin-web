package http

import (
	"fmt"
	"github.com/CoderSamYhc/gin-web/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Message: "success",
		Data: data,
	})
	c.Abort()
}

func ErrorResponse(c *gin.Context, code int, data ...interface{}) {
	val, ok := config.CodeMsgMap[code]
	if !ok {
		code = 99999
	}
	msg := fmt.Sprintf(val, data...)
	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: msg,
		Data: map[string]interface{}{},
	})
	c.Abort()
}
