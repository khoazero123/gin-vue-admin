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

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func CreateAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := service.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Any("err", err))
		response.FailWithMessage("Creation failed"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "Create success", c)
	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func CopyAuthority(c *gin.Context) {
	var copyInfo response.SysAuthorityCopyResponse
	_ = c.ShouldBindJSON(&copyInfo)
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := service.CopyAuthority(copyInfo); err != nil {
		global.GVA_LOG.Error("Copy failure!", zap.Any("err", err))
		response.FailWithMessage("Copy failure"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "Copy success", c)
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /authority/deleteAuthority [post]
func DeleteAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteAuthority(&authority); err != nil { // Remove the role before you need to judge whether the user is using this role
		global.GVA_LOG.Error("Failed to delete!", zap.Any("err", err))
		response.FailWithMessage("Failed to delete"+err.Error(), c)
	} else {
		response.OkWithMessage("Successfully deleted", c)
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func UpdateAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authority := service.UpdateAuthority(auth); err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Any("err", err))
		response.FailWithMessage("Update failed"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authority}, "Update completed", c)
	}
}

// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetAuthorityInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("Acquisition failure!", zap.Any("err", err))
		response.FailWithMessage("Acquisition failure"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Get successful", c)
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func SetDataAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.SetDataAuthority(auth); err != nil {
		global.GVA_LOG.Error("Setup failed!", zap.Any("err", err))
		response.FailWithMessage("Setup failed"+err.Error(), c)
	} else {
		response.OkWithMessage("Set success", c)
	}
}
