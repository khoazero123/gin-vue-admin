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

// @Tags SysDictionaryDetail
// @Summary 创建SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func CreateSysDictionaryDetail(c *gin.Context) {
	var detail model.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := service.CreateSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Any("err", err))
		response.FailWithMessage("Creation failed", c)
	} else {
		response.OkWithMessage("Create success", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 删除SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func DeleteSysDictionaryDetail(c *gin.Context) {
	var detail model.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := service.DeleteSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("Failed to delete!", zap.Any("err", err))
		response.FailWithMessage("Failed to delete", c)
	} else {
		response.OkWithMessage("Successfully deleted", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 更新SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "更新SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func UpdateSysDictionaryDetail(c *gin.Context) {
	var detail model.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := service.UpdateSysDictionaryDetail(&detail); err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Any("err", err))
		response.FailWithMessage("Update failed", c)
	} else {
		response.OkWithMessage("Update completed", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 用id查询SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysDictionaryDetail true "用id查询SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func FindSysDictionaryDetail(c *gin.Context) {
	var detail model.SysDictionaryDetail
	_ = c.ShouldBindQuery(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysDictionaryDetail := service.GetSysDictionaryDetail(detail.ID); err != nil {
		global.GVA_LOG.Error("Query failed!", zap.Any("err", err))
		response.FailWithMessage("Query failed", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionaryDetail": resysDictionaryDetail}, "search successful", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 分页获取SysDictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysDictionaryDetailSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func GetSysDictionaryDetailList(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSysDictionaryDetailInfoList(pageInfo); err != nil {
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
