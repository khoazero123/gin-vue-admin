package model

import (
	"gin-vue-admin/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:User UUID"`                                                            // User UUID
	Username    string       `json:"userName" gorm:"comment:User login name"`                                                  // User login name
	Password    string       `json:"-"  gorm:"comment:User login password"`                                                    // User login password
	NickName    string       `json:"nickName" gorm:"default:system user;comment:User's Nickname"`                              // User's Nickname
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:profile picture"` // profile picture
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:User role"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:User role ID"` // User role ID
}
