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

// @Tags {{.StructName}}
// @Summary Create {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Create {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func Create{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	if err := service.Create{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("Creation failed!", zap.Any("err", err))
		response.FailWithMessage("Creation failed", c)
	} else {
		response.OkWithMessage("Create success", c)
	}
}

// @Tags {{.StructName}}
// @Summary Delete {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Delete {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func Delete{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	if err := service.Delete{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("Failed to delete!", zap.Any("err", err))
		response.FailWithMessage("Failed to delete", c)
	} else {
		response.OkWithMessage("Successfully deleted", c)
	}
}

// @Tags {{.StructName}}
// @Summary Batch deletion {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Batch deletion {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Batch delete success"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}}ByIds [delete]
func Delete{{.StructName}}ByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.Delete{{.StructName}}ByIds(IDS); err != nil {
        global.GVA_LOG.Error("Batch deletion failed!", zap.Any("err", err))
		response.FailWithMessage("Batch deletion failed", c)
	} else {
		response.OkWithMessage("Batch delete success", c)
	}
}

// @Tags {{.StructName}}
// @Summary Update {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Update {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Update completed"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
func Update{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindJSON(&{{.Abbreviation}})
	if err := service.Update{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("Update failed!", zap.Any("err", err))
		response.FailWithMessage("Update failed", c)
	} else {
		response.OkWithMessage("Update completed", c)
	}
}

// @Tags {{.StructName}}
// @Summary Id query {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Id query {{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
func Find{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} model.{{.StructName}}
	_ = c.ShouldBindQuery(&{{.Abbreviation}})
	if err, re{{.Abbreviation}} := service.Get{{.StructName}}({{.Abbreviation}}.ID); err != nil {
        global.GVA_LOG.Error("Query failed!", zap.Any("err", err))
		response.FailWithMessage("Query failed", c)
	} else {
		response.OkWithData(gin.H{"re{{.Abbreviation}}": re{{.Abbreviation}}}, c)
	}
}

// @Tags {{.StructName}}
// @Summary Page acquisition {{.StructName}} List
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.{{.StructName}}Search true "Page acquisition {{.StructName}} List"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func Get{{.StructName}}List(c *gin.Context) {
	var pageInfo request.{{.StructName}}Search
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.Get{{.StructName}}InfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("Acquisition failure", zap.Any("err", err))
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
