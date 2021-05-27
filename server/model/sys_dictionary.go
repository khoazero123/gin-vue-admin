// Automatic generation template SysDictionary
package model

import (
	"gin-vue-admin/global"
)

// If you contain Time.Time, please ask your own Import Time package.
type SysDictionary struct {
	global.GVA_MODEL
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:Dictionary name (middle)"`  // 字典名（中）
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:Dictionary name (English)"` // 字典名（英）
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:status"`              // 状态
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:description"`               // 描述
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}
