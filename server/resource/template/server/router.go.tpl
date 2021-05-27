package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.StructName}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	{
		{{.StructName}}Router.POST("create{{.StructName}}", v1.Create{{.StructName}})   // New construction{{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}", v1.Delete{{.StructName}}) // Delete {{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}ByIds", v1.Delete{{.StructName}}ByIds) // Batch deletion{{.StructName}}
		{{.StructName}}Router.PUT("update{{.StructName}}", v1.Update{{.StructName}})    // Update {{.StructName}}
		{{.StructName}}Router.GET("find{{.StructName}}", v1.Find{{.StructName}})        // Acquisition according to ID {{.StructName}}
		{{.StructName}}Router.GET("get{{.StructName}}List", v1.Get{{.StructName}}List)  // Obtain {{.StructName}}List
	}
}
