package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(c *gin.Context, status int, data any, message string) {
	c.JSON(http.StatusOK, map[string]any{
		"data":    data,
		"message": message,
		"code":    status,
	})
}

func ResponseSuccess(c *gin.Context, data any) {
	response(c, 200, data, "Success")
}

func ResponseFail(c *gin.Context, status int, data any, message string) {
	response(c, status, data, message)
}
