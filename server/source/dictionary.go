package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Dictionary = new(dictionary)

type dictionary struct{}

var status = new(bool)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_dictionaries Table data initialization
func (d *dictionary) Init() error {
	*status = true
	var dictionaries = []model.SysDictionary{
		{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "gender", Type: "sex", Status: status, Desc: "Gender dictionary"},
		{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "Database INT type", Type: "int", Status: status, Desc: "INT type corresponding database type"},
		{GVA_MODEL: global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "Database time date type", Type: "time.Time", Status: status, Desc: "Database time date type"},
		{GVA_MODEL: global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "Database floating point", Type: "float64", Status: status, Desc: "Database floating point"},
		{GVA_MODEL: global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "Database string", Type: "string", Status: status, Desc: "Database string"},
		{GVA_MODEL: global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "Database BOOL Type", Type: "bool", Status: status, Desc: "Database BOOL Type"},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]model.SysDictionary{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_dictionaries Table initial data already exists!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // Magnification
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_dictionaries Table initial data success!")
		return nil
	})
}
