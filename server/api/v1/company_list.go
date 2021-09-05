package v1

import (
	
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags company
// @Summary
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"查询成功"}"
// @Router /json/companyList [get]
func CompanyList( c *gin.Context) {
	if err := global.GVA_DB.Migrator().HasTable("company_List"); !err {
		service.CreateCompanyTable()
	}
	err, list := service.GetCompanyList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(list, "查询成功", c)
	}
}


// @Tags company
// @Summary
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"查询成功"}"
// @Router /json/addCompanyList [Post]
func AddCompanyList( c *gin.Context) {
	var list model.CompanyList;
	_ = c.ShouldBindJSON(&list)
	if err := service.AddCompanyList(list); err != nil {
		response.FailWithMessage(err.Error(), c)
	}else {
		response.OkWithMessage("操作成功", c)
	}
}

// @Tags company
// @Summary
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /json/companyList [delete]
func DeleteCompanylist(c *gin.Context) {
	var company model.CompanyList
	_ = c.ShouldBindJSON(&company)
	if err := service.DeleteCompanyList(company); err != nil {
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags company
// @Summary
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /json/companyList [post]
func UpdateCompanylist(c *gin.Context) {
	var company model.CompanyList
	_ = c.ShouldBindJSON(&company)
	if err := service.UpdateCompanylist(company); err != nil {
		response.FailWithMessage("更新失败!", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}