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

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createApi [post]
func CreateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateApi(api); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Any("err", err))
		response.FailWithMessage("Creation failed", c)
	} else {
		response.OkWithMessage("Create success", c)
	}
}

// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /api/deleteApi [post]
func DeleteApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteApi(api); err != nil {
		global.GVA_LOG.Error("Failed to delete!", zap.Any("err", err))
		response.FailWithMessage("Failed to delete", c)
	} else {
		response.OkWithMessage("Successfully deleted", c)
	}
}

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(c *gin.Context) {
	var pageInfo request.SearchApiParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
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

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, api := service.GetApiById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failure!", zap.Any("err", err))
		response.FailWithMessage("Acquisition failure", c)
	} else {
		response.OkWithData(response.SysAPIResponse{Api: api}, c)
	}
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/updateApi [post]
func UpdateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateApi(api); err != nil {
		global.GVA_LOG.Error("Fail to edit!", zap.Any("err", err))
		response.FailWithMessage("Fail to edit", c)
	} else {
		response.OkWithMessage("Successfully modified", c)
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(c *gin.Context) {
	if err, apis := service.GetAllApis(); err != nil {
		global.GVA_LOG.Error("Acquisition failure!", zap.Any("err", err))
		response.FailWithMessage("Acquisition failure", c)
	} else {
		response.OkWithDetailed(response.SysAPIListResponse{Apis: apis}, "Get successful", c)
	}
}

// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /api/deleteApisByIds [delete]
func DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := service.DeleteApisByIds(ids); err != nil {
		global.GVA_LOG.Error("Failed to delete!", zap.Any("err", err))
		response.FailWithMessage("Failed to delete", c)
	} else {
		response.OkWithMessage("Successfully deleted", c)
	}
}
