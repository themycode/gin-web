package handleFuns

import (
	"github.com/CoderSamYhc/gin-web/http"
	"github.com/gin-gonic/gin"
)


func Index (c *gin.Context){
	http.SuccessResponse(c, nil)
}
