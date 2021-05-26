package model

import (
	"gin-vue-admin/global"
)

type SysBaseMenu struct {
	global.GVA_MODEL
	MenuLevel     uint                                              `json:"-"`
	ParentId      string                                            `json:"parentId" gorm:"comment:Parent Menu ID"`                     // 父菜单ID
	Path          string                                            `json:"path" gorm:"comment:Routing Path"`                           // 路由path
	Name          string                                            `json:"name" gorm:"comment:Route Name"`                             // 路由name
	Hidden        bool                                              `json:"hidden" gorm:"comment:Is it hidden in the list?"`            // 是否在列表隐藏
	Component     string                                            `json:"component" gorm:"comment:Corresponding front end file path"` // 对应前端文件路径
	Sort          int                                               `json:"sort" gorm:"comment:Sort tag"`                               // 排序标记
	Meta          `json:"meta" gorm:"comment:Additional attribute"` // 附加属性
	SysAuthoritys []SysAuthority                                    `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu                                     `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter                            `json:"parameters"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:Whether to cache"`                                   // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:Whether it is the basic routing (in development)"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:Menu name"`                                              // 菜单名
	Icon        string `json:"icon" gorm:"comment:Menu icon"`                                               // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:Automatically turn off TAB"`                          // 自动关闭tab
}

type SysBaseMenuParameter struct {
	global.GVA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:Address bar carries parameters to params or query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"comment:KEY with the address bar"`                           // 地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:Address bar carrying the value of parameters"`     // 地址栏携带参数的值
}
