package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	// "gin-vue-admin/global"
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