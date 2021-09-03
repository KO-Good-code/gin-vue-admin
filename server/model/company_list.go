package model

import (
	"gin-vue-admin/global"
)

type CompanyList struct {
	global.GVA_MODEL
	Name	string
	Num 	int
}

func (CompanyList) TableName() string {
	return "company_list"
}