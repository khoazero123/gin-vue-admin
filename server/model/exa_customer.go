package model

import (
	"gin-vue-admin/global"
)

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName       string  `json:"customerName" form:"customerName" gorm:"comment:Customer name"`                          // Customer name
	CustomerPhoneData  string  `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:Customer mobile phone number"` // Customer mobile phone number
	SysUserID          uint    `json:"sysUserId" form:"sysUserId" gorm:"comment:Management ID"`                                // Management ID
	SysUserAuthorityID string  `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:Management role ID"`         // Management role ID
	SysUser            SysUser `json:"sysUser" form:"sysUser" gorm:"comment:Management details"`                               // Management details
}
