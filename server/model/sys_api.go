package model

import (
	"gin-vue-admin/global"
)

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"comment:API path"`                       // API path
	Description string `json:"description" gorm:"comment:API Chinese description"` // API Chinese description
	ApiGroup    string `json:"apiGroup" gorm:"comment:API"`                        // API
	Method      string `json:"method" gorm:"default:POST" gorm:"comment:method"`   // method:创建POST(默认)|查看GET|更新PUT|删除DELETE
}
