package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExaCustomer
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户用户名, 客户手机号码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customer/customer [post]
func CreateExaCustomer(c *gin.Context) {
	var customer model.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customer.SysUserID = getUserID(c)
	customer.SysUserAuthorityID = getUserAuthorityId(c)
	if err := service.CreateExaCustomer(customer); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Any("err", err))
		response.FailWithMessage("Creation failed", c)
	} else {
		response.OkWithMessage("Create success", c)
	}
}

// @Tags ExaCustomer
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /customer/customer [delete]
func DeleteExaCustomer(c *gin.Context) {
	var customer model.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteExaCustomer(customer); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Any("err", err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// @Tags ExaCustomer
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户ID, 客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/customer [put]
func UpdateExaCustomer(c *gin.Context) {
	var customer model.ExaCustomer
	_ = c.ShouldBindJSON(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(customer, utils.CustomerVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateExaCustomer(&customer); err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Any("err", err))
		response.FailWithMessage("Update failed!", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// @Tags ExaCustomer
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "客户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func GetExaCustomer(c *gin.Context) {
	var customer model.ExaCustomer
	_ = c.ShouldBindQuery(&customer)
	if err := utils.Verify(customer.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, data := service.GetExaCustomer(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failure!", zap.Any("err", err))
		response.FailWithMessage("Acquisition failure", c)
	} else {
		response.OkWithDetailed(response.ExaCustomerResponse{Customer: data}, "Get successful", c)
	}
}

// @Tags ExaCustomer
// @Summary 分页获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
func GetExaCustomerList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, customerList, total := service.GetCustomerInfoList(getUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failure!", zap.Any("err", err))
		response.FailWithMessage(fmt.Sprintf("Get failed: %v", err), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Get successful", c)
	}
}
