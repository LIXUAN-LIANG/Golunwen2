package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Success(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Fail(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusBadRequest, 400, data, msg)

}
