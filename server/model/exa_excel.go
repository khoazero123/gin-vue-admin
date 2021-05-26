package model

type ExcelInfo struct {
	FileName string        `json:"fileName"` // file name
	InfoList []SysBaseMenu `json:"infoList"`
}
