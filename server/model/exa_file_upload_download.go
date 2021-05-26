package model

import (
	"gin-vue-admin/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"comment:file name"`   // file name
	Url  string `json:"url" gorm:"comment:File address"` // File address
	Tag  string `json:"tag" gorm:"comment:File label"`   // File label
	Key  string `json:"key" gorm:"comment:Numbering"`    // Numbering
}
