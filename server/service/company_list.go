package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

/**
//@function: CreateCompanyTable
//@description: 创建公司数据表
//@param: list model.CompanyList
//@return: model.CompanyList
*/
func CreateCompanyTable()(err error) {
	// var companyList model.CompanyList
	err = global.GVA_DB.Migrator().CreateTable(&model.CompanyList{});
	return err
}


/**
//@function: GetCompanyList
//@description: 获取公司数据
//@param: list model.CompanyList
//@return: model.CompanyList
*/
func GetCompanyList() (err error, list model.CompanyList) {
	err = global.GVA_DB.Find(&list).Error
	return err, list
}