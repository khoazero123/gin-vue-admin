package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// @Tags InitDB
// @Summary Initialize user database
// @Produce  application/json
// @Param data body request.InitDB true "Initialization database parameters"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Automatically create database success"}"
// @Router /init/initdb [post]
func InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("Unauthorized access")
		response.FailWithMessage("Unauthorized access", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("Parameter verification", zap.Any("err", err))
		response.FailWithMessage("Parameter verification", c)
		return
	}
	if err := service.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("Automatically create database failed", zap.Any("err", err))
		response.FailWithMessage("Automatically create a database failed, please see the background log", c)
		return
	}
	response.OkWithData("Automatically create database success", c)
}

// @Tags CheckDB
// @Summary Initialize user database
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"Detection completion"}"
// @Router /init/checkdb [post]
func CheckDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Info("Database does not need to initialize")
		response.OkWithDetailed(gin.H{
			"needInit": false,
		}, "Database does not need to initialize", c)
		return
	} else {
		global.GVA_LOG.Info("Go to the initialization database")
		response.OkWithDetailed(gin.H{
			"needInit": true,
		}, "Go to the initialization database", c)
		return
	}
}
