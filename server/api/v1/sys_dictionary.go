package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "SysDictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDictionary/createSysDictionary [post]
func CreateSysDictionary(c *gin.Context) {
	var dictionary model.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := service.CreateSysDictionary(dictionary); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Any("err", err))
		response.FailWithMessage("Creation failed", c)
	} else {
		response.OkWithMessage("Create success", c)
	}
}

// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "SysDictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /sysDictionary/deleteSysDictionary [delete]
func DeleteSysDictionary(c *gin.Context) {
	var dictionary model.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := service.DeleteSysDictionary(dictionary); err != nil {
		global.GVA_LOG.Error("Failed to delete!", zap.Any("err", err))
		response.FailWithMessage("Failed to delete", c)
	} else {
		response.OkWithMessage("Successfully deleted", c)
	}
}

// @Tags SysDictionary
// @Summary 更新SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "SysDictionary模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionary/updateSysDictionary [put]
func UpdateSysDictionary(c *gin.Context) {
	var dictionary model.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := service.UpdateSysDictionary(&dictionary); err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Any("err", err))
		response.FailWithMessage("Update failed", c)
	} else {
		response.OkWithMessage("Update completed", c)
	}
}

// @Tags SysDictionary
// @Summary 用id查询SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionary true "ID或字典英名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionary/findSysDictionary [get]
func FindSysDictionary(c *gin.Context) {
	var dictionary model.SysDictionary
	_ = c.ShouldBindQuery(&dictionary)
	if err, sysDictionary := service.GetSysDictionary(dictionary.Type, dictionary.ID); err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Any("err", err))
		response.FailWithMessage("Query failed", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "search successful", c)
	}
}

// @Tags SysDictionary
// @Summary 分页获取SysDictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysDictionarySearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/getSysDictionaryList [get]
func GetSysDictionaryList(c *gin.Context) {
	var pageInfo request.SysDictionarySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetSysDictionaryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("Acquisition failure!", zap.Any("err", err))
		response.FailWithMessage("Acquisition failure", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Get successful", c)
	}
}
