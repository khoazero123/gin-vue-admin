package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.SysUser{
	{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "Super administrator", HeaderImg: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
	{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", NickName: "QMPlusUser", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users Table data initialization
func (a *admin) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users The initial data of the table already exists!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // Magnification
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users Table initial data success!")
		return nil
	})
}
