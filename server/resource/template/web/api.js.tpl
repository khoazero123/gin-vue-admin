import service from '@/utils/request'

// @Tags {{.StructName}}
// @Summary create{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "create{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
export const create{{.StructName}} = (data) => {
     return service({
         url: "/{{.Abbreviation}}/create{{.StructName}}",
         method: 'post',
         data
     })
 }


// @Tags {{.StructName}}
// @Summary delete{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "delete{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
 export const delete{{.StructName}} = (data) => {
     return service({
         url: "/{{.Abbreviation}}/delete{{.StructName}}",
         method: 'delete',
         data
     })
 }

// @Tags {{.StructName}}
// @Summary delete{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "batch deletion{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successfully deleted"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
 export const delete{{.StructName}}ByIds = (data) => {
     return service({
         url: "/{{.Abbreviation}}/delete{{.StructName}}ByIds",
         method: 'delete',
         data
     })
 }

// @Tags {{.StructName}}
// @Summary Update{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Update{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
 export const update{{.StructName}} = (data) => {
     return service({
         url: "/{{.Abbreviation}}/update{{.StructName}}",
         method: 'put',
         data
     })
 }


// @Tags {{.StructName}}
// @Summary Id query {{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "Id query{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"search successful"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
 export const find{{.StructName}} = (params) => {
     return service({
         url: "/{{.Abbreviation}}/find{{.StructName}}",
         method: 'get',
         params
     })
 }


// @Tags {{.StructName}}
// @Summary Page acquisition{{.StructName}}List
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "Page acquisition {{.StructName}}List"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
 export const get{{.StructName}}List = (params) => {
     return service({
         url: "/{{.Abbreviation}}/get{{.StructName}}List",
         method: 'get',
         params
     })
 }