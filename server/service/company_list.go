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
func GetCompanyList() (err error, list[] model.CompanyList) {
	err = global.GVA_DB.Find(&list).Error
	return err, list
}

/**
//@function: GetCompanyList
//@description: 增加公司数据
//@param: list model.CompanyList
//@return: model.CompanyList
*/
func AddCompanyList(e model.CompanyList) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

/**
//@function: GetCompanyList
//@description: 删除公司数据
//@param: list model.CompanyList
//@return: model.CompanyList
*/
func DeleteCompanyList(e model.CompanyList) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func UpdateCompanylist(e model.CompanyList) (err error){
	err = global.GVA_DB.Save(e).Error
	return err
}