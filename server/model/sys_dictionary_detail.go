// 自动生成模板SysDictionaryDetail
package model

import (
	"gin-vue-admin/global"
)

// If you contain Time.Time, please ask your own Import Time package.
type SysDictionaryDetail struct {
	global.GVA_MODEL
	Label           string `json:"label" form:"label" gorm:"column:label;comment:Display value"`                                  // 展示值
	Value           int    `json:"value" form:"value" gorm:"column:value;comment:Dictionary value"`                               // 字典值
	Status          *bool  `json:"status" form:"status" gorm:"column:status;comment:Enable state"`                                // 启用状态
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:Sort tag"`                                          // 排序标记
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id;comment:Associated tag"` // 关联标记
}
