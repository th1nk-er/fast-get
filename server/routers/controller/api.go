package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/th1nk-er/fast-get/server/core"
	"github.com/th1nk-er/fast-get/server/pkg/utils"
	"net/http"
)

type MsgController struct{}

func NewMsgController() *MsgController {
	return &MsgController{}
}

func (m MsgController) GetMessage(c *gin.Context) {
	var model core.GetModel
	if c.ShouldBind(&model) != nil || model.MsgID == "" {
		utils.ResponseFail(c, 1001, "", "Parameter error")
		return
	}
	code, msg, err := core.GetMessage(model)
	if model.Raw {
		if err != nil {

			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.String(http.StatusOK, msg)
		return
	}
	if err != nil {
		utils.ResponseFail(c, code, nil, err.Error())
		return
	}
	utils.ResponseSuccess(c, msg)
}

func (m MsgController) SaveMessage(c *gin.Context) {
	var model core.SaveModel
	if c.ShouldBind(&model) != nil {
		utils.ResponseFail(c, 1001, "", "Parameter error")
		return
	}
	code, msgID, err := core.SaveMessage(model)
	if err != nil {
		utils.ResponseFail(c, code, nil, err.Error())
		return
	}
	utils.ResponseSuccess(c, msgID)
}

func (m MsgController) EditMessage(c *gin.Context) {
	var model core.EditModel
	if c.ShouldBind(&model) != nil {
		utils.ResponseFail(c, 1001, "", "Parameter error")
		return
	}
	code, err := core.EditMessage(model)
	if err != nil {
		utils.ResponseFail(c, code, nil, err.Error())
		return
	}
	utils.ResponseSuccess(c, nil)
}
