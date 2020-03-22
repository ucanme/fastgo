package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ucanme/fastgo/controller/response"
)

func Demo(c *gin.Context) {
	response.Success(c, map[string]interface{}{"Hello": "World"})
	return
}
