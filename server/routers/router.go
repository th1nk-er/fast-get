package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	msgRouter(r)
	return r
}
