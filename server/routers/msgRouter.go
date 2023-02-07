package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/th1nk-er/fast-get/server/routers/controller"
)

func msgRouter(e *gin.Engine) {
	msgController := controller.NewMsgController()
	e.GET("/msg", msgController.GetMessage)
	e.POST("/msg", msgController.SaveMessage)
	e.PUT("/msg", msgController.EditMessage)
}
