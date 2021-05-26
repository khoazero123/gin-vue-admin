// Automatic generation template SYSOPERATIONRecord
package model

import (
	"gin-vue-admin/global"
	"time"
)

// If you contain Time.Time, please ask your own Import Time package.
type SysOperationRecord struct {
	global.GVA_MODEL
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:Request IP"`                                     // Request IP
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:Request method"`                     // Request method
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:Request path"`                             // Request path
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:Request status"`                     // Request status
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:delay" swaggertype:"string"`      // delay
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:proxy"`                                 // proxy
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:Error message"` // Error message
	Body         string        `json:"body" form:"body" gorm:"type:longtext;column:body;comment:Request Body"`               // Request Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:longtext;column:resp;comment:Response body"`              // Response body
	UserID       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:User ID"`                         // User ID
	User         SysUser       `json:"user"`
}
