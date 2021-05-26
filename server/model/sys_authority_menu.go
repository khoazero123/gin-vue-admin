package model

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:Menu ID"`
	AuthorityId string                 `json:"-" gorm:"comment:Role ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}
