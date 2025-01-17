package model

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time      // Create time
	UpdatedAt       time.Time      // Update time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:Role ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:character name"`                            // 角色名
	ParentId        string         `json:"parentId" gorm:"comment:Father ID"`                                      // 父角色ID
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter   string         `json:"defaultRouter" gorm:"comment:Default menu;default:dashboard"` // 默认菜单(默认dashboard)
}
