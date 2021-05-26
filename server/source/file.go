package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var File = new(file)

type file struct{}

var files = []model.ExaFileUploadAndDownload{
	{global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "http://qmplusimg.henrongyi.top/gvalogo.png", "png", "158787308910.png"},
	{global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", "png", "1587973709logo.png"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: exa_file_upload_and_downloads Table initialization data
func (f *file) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.ExaFileUploadAndDownload{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> exa_file_upload_and_downloads Table initial data already exists!")
			return nil
		}
		if err := tx.Create(&files).Error; err != nil { // Magnification
			return err
		}
		color.Info.Println("\n[Mysql] --> exa_file_upload_and_downloads Table initial data success!")
		return nil
	})
}
