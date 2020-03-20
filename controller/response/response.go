package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/ucanme/fastgo/library/log"
)

func Response(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, map[string]interface{}{
		"data":       data,
		"error_no":   code,
		"error_msg":  msg,
		"request_id": log.GetRequestID(c),
	})
}

func Fail(c *gin.Context, code int, msg string) {
	Response(c, code, msg, nil)
}

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	msg := "success"
	fmt.Println("Data", data)
	Response(c, 0, msg, data)
}
