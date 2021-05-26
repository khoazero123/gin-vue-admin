package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []model.SysApi{
	{global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", "User login", "base", "POST"},
	{global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/register", "User registration", "user", "POST"},
	{global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/createApi", "Create an API", "api", "POST"},
	{global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiList", "Get API list", "api", "POST"},
	{global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiById", "Get API details", "api", "POST"},
	{global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApi", "Delete API", "api", "POST"},
	{global.GVA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/updateApi", "Update API", "api", "POST"},
	{global.GVA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getAllApis", "Get all APIs", "api", "POST"},
	{global.GVA_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/createAuthority", "Creating a Role", "authority", "POST"},
	{global.GVA_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/deleteAuthority", "Delete role", "authority", "POST"},
	{global.GVA_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/getAuthorityList", "Get a list of roles", "authority", "POST"},
	{global.GVA_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenu", "Get menu tree", "menu", "POST"},
	{global.GVA_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuList", "Page Get Basics Menu List", "menu", "POST"},
	{global.GVA_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addBaseMenu", "Add menu", "menu", "POST"},
	{global.GVA_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuTree", "Get user dynamic routes", "menu", "POST"},
	{global.GVA_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addMenuAuthority", "Increase Menu and Role Association", "menu", "POST"},
	{global.GVA_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuAuthority", "Get specified role Menu", "menu", "POST"},
	{global.GVA_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/deleteBaseMenu", "Delete menute menu", "menu", "POST"},
	{global.GVA_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/updateBaseMenu", "Update menu", "menu", "POST"},
	{global.GVA_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuById", "Get menu according to ID", "menu", "POST"},
	{global.GVA_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/changePassword", "change Password", "user", "POST"},
	{global.GVA_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/getUserList", "Get user list", "user", "POST"},
	{global.GVA_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserAuthority", "Modify user role", "user", "POST"},
	{global.GVA_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/upload", "File upload", "fileUploadAndDownload", "POST"},
	{global.GVA_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/getFileList", "Get list of uploaded files", "fileUploadAndDownload", "POST"},
	{global.GVA_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/updateCasbin", "Change role API permissions", "casbin", "POST"},
	{global.GVA_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/getPolicyPathByAuthorityId", "Get permission list", "casbin", "POST"},
	{global.GVA_MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/deleteFile", "Delete Files", "fileUploadAndDownload", "POST"},
	{global.GVA_MODEL{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/jwt/jsonInBlacklist", "JWT joins blacklist(drop out)", "jwt", "POST"},
	{global.GVA_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setDataAuthority", "Set role resource permissions", "authority", "POST"},
	{global.GVA_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getSystemConfig", "Get profile content", "system", "POST"},
	{global.GVA_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/setSystemConfig", "Set the configuration file content", "system", "POST"},
	{global.GVA_MODEL{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "Create a customer", "customer", "POST"},
	{global.GVA_MODEL{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "Update customer", "customer", "PUT"},
	{global.GVA_MODEL{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "Delete customers", "customer", "DELETE"},
	{global.GVA_MODEL{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "Get a single customer", "customer", "GET"},
	{global.GVA_MODEL{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customerList", "Get a list of customers", "customer", "GET"},
	{global.GVA_MODEL{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/casbinTest/:pathParam", "RESTful mode test", "casbin", "GET"},
	{global.GVA_MODEL{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/createTemp", "Automation code", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/updateAuthority", "Update role information", "authority", "PUT"},
	{global.GVA_MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/copyAuthority", "Copy role", "authority", "POST"},
	{global.GVA_MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/deleteUser", "delete users", "user", "DELETE"},
	{global.GVA_MODEL{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/createSysDictionaryDetail", "New dictionary content", "sysDictionaryDetail", "POST"},
	{global.GVA_MODEL{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/deleteSysDictionaryDetail", "Delete dictionary content", "sysDictionaryDetail", "DELETE"},
	{global.GVA_MODEL{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/updateSysDictionaryDetail", "Update dictionary content", "sysDictionaryDetail", "PUT"},
	{global.GVA_MODEL{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/findSysDictionaryDetail", "Get dictionary according to ID", "sysDictionaryDetail", "GET"},
	{global.GVA_MODEL{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/getSysDictionaryDetailList", "Get a list of dictionary contents", "sysDictionaryDetail", "GET"},
	{global.GVA_MODEL{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/createSysDictionary", "New dictionary", "sysDictionary", "POST"},
	{global.GVA_MODEL{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/deleteSysDictionary", "Delete dictionary", "sysDictionary", "DELETE"},
	{global.GVA_MODEL{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/updateSysDictionary", "Update dictionary", "sysDictionary", "PUT"},
	{global.GVA_MODEL{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/findSysDictionary", "Get a dictionary according to ID", "sysDictionary", "GET"},
	{global.GVA_MODEL{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/getSysDictionaryList", "Get a dictionary list", "sysDictionary", "GET"},
	{global.GVA_MODEL{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/createSysOperationRecord", "New operation record", "sysOperationRecord", "POST"},
	{global.GVA_MODEL{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecord", "Delete operation record", "sysOperationRecord", "DELETE"},
	{global.GVA_MODEL{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/findSysOperationRecord", "Operation record according to ID", "sysOperationRecord", "GET"},
	{global.GVA_MODEL{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/getSysOperationRecordList", "Get a list of operation records", "sysOperationRecord", "GET"},
	{global.GVA_MODEL{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getTables", "Get database tables", "autoCode", "GET"},
	{global.GVA_MODEL{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getDB", "Get all databases", "autoCode", "GET"},
	{global.GVA_MODEL{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getColumn", "Get all the fields of the selected table", "autoCode", "GET"},
	{global.GVA_MODEL{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecordByIds", "Batch delete operation history", "sysOperationRecord", "DELETE"},
	{global.GVA_MODEL{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/upload", "Plug-in version slice upload", "simpleUploader", "POST"},
	{global.GVA_MODEL{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/checkFileMd5", "Document completion verification", "simpleUploader", "GET"},
	{global.GVA_MODEL{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/mergeFileMd5", "Upload finishing files", "simpleUploader", "GET"},
	{global.GVA_MODEL{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserInfo", "Set user information", "user", "PUT"},
	{global.GVA_MODEL{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getServerInfo", "Get server information", "system", "POST"},
	{global.GVA_MODEL{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/email/emailTest", "Send test mail", "email", "POST"},
	{global.GVA_MODEL{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/preview", "Preview Automation Code", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/importExcel", "Import Excel", "excel", "POST"},
	{global.GVA_MODEL{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/loadExcel", "Download Excel", "excel", "GET"},
	{global.GVA_MODEL{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/exportExcel", "Export Excel", "excel", "POST"},
	{global.GVA_MODEL{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/downloadTemplate", "Download Excel Template", "excel", "GET"},
	{global.GVA_MODEL{ID: 85, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApisByIds", "Batch Delete API", "api", "DELETE"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_apis Table data initialization
func (a *api) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]model.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis The initial data of the table already exists!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // Magnification
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis Table initial data success!")
		return nil
	})
}
