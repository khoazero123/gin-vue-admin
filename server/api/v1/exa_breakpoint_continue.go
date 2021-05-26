package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("Receive file failed!", zap.Any("err", err))
		response.FailWithMessage("Receive file failed", c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.GVA_LOG.Error("File read failed!", zap.Any("err", err))
		response.FailWithMessage("File read failed", c)
		return
	}
	defer f.Close()
	cen, _ := ioutil.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.GVA_LOG.Error("Check MD5 failure!", zap.Any("err", err))
		response.FailWithMessage("Check MD5 failure", c)
		return
	}
	err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("Find or create a record failed!", zap.Any("err", err))
		response.FailWithMessage("Find or create a record failed", c)
		return
	}
	err, pathc := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("Breakpoint renewal failed!", zap.Any("err", err))
		response.FailWithMessage("Breakpoint renewal failed", c)
		return
	}

	if err = service.CreateFileChunk(file.ID, pathc, chunkNumber); err != nil {
		global.GVA_LOG.Error("Create file log failed!", zap.Any("err", err))
		response.FailWithMessage("Create file log failed", c)
		return
	}
	response.OkWithMessage("Slice creation success", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("Find failed!", zap.Any("err", err))
		response.FailWithMessage("Find failed", c)
	} else {
		response.OkWithDetailed(response.FileResponse{File: file}, "Find success", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 创建文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("File creation failed!", zap.Any("err", err))
		response.FailWithDetailed(response.FilePathResponse{FilePath: filePath}, "File creation failed", c)
	} else {
		response.OkWithDetailed(response.FilePathResponse{FilePath: filePath}, "Document creation success", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"缓存切片删除成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func RemoveChunk(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath := c.Query("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		global.GVA_LOG.Error("Cache slice delete failed!", zap.Any("err", err))
		response.FailWithDetailed(response.FilePathResponse{FilePath: filePath}, "Cache slice delete failed", c)
	} else {
		response.OkWithDetailed(response.FilePathResponse{FilePath: filePath}, "Cache slice delete success", c)
	}
}
