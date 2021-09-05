package model

import (
	"gin-vue-admin/global"
)

type CompanyList struct {
	global.GVA_MODEL
	Name	string `json:"name" form:"name" gorm:"comment:公司名称"` 
	Num 	int `json:"num" form:"num" gorm:"comment:公司投标价"` 
}

func (CompanyList) TableName() string {
	return "company_list"
}