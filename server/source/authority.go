package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

var authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "general user", ParentId: "0", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "Ordinary subsidiary", ParentId: "888", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "Test role", ParentId: "0", DefaultRouter: "dashboard"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authorities Table data initialization
func (a *authority) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_authorities The initial data of the table already exists!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // Magnification
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authorities Table initial data success!")
		return nil
	})
}
