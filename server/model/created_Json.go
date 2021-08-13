package model

import (
	"gin-vue-admin/global"
)

// 表单数据， 切片结构体
type TableMap struct {
	global.GVA_MODEL
	ID uint `json:"id" form:"id"`
	Label string `json:"label" form:"label"`
	Key string `json:"Key" form:"Key"`
	DataType string `json:"dataType" form:"dataType"`
	DefaultValue string `json:"defaultValue" form:"defaultValue"`
}

// 数据表单结构体
type Table struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" require`
	Data []TableMap `json:"data" form:"data"`
}