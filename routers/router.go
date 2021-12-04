package routers

import (
	"github.com/CoderSamYhc/gin-web/http/handleFuns"
	user_http "github.com/CoderSamYhc/gin-web/http/handleFuns/user"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		handleFuns.Index(c)
	})

	v1 := r.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, "v1")
		})
		v1.GET("/read", func(c *gin.Context) {
			c.JSON(200, "read")
		})
		user := v1.Group("user")
		{
			user.POST("/add", func(c *gin.Context) {
				user_http.Add(c)
			})
			user.POST("/find", func(c *gin.Context) {
				user_http.Find(c)
			})
		}
	}
}
