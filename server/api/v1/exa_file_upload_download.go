package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func UploadFile(c *gin.Context) {
	var file model.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("Receive file failed!", zap.Any("err", err))
		response.FailWithMessage("Receive file failed", c)
		return
	}
	err, file = service.UploadFile(header, noSave) // After the file is uploaded, get the file path
	if err != nil {
		global.GVA_LOG.Error("Modify the database link failed!", zap.Any("err", err))
		response.FailWithMessage("Modify the database link failed", c)
		return
	}
	response.OkWithDetailed(response.ExaFileResponse{File: file}, "Upload success", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary Delete Files
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "Passing the ID in the file"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func DeleteFile(c *gin.Context) {
	var file model.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := service.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("Failed to delete!", zap.Any("err", err))
		response.FailWithMessage("Failed to delete", c)
		return
	}
	response.OkWithMessage("Successfully deleted", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary Paging file list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "Page number, per page size"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /fileUploadAndDownload/getFileList [post]
func GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetFileRecordInfoList(pageInfo)
	if err != nil {
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
