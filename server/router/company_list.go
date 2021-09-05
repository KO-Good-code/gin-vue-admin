package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCreatdJson(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("json").Use(middleware.OperationRecord())
	{
		ApiRouter.GET("companyList", v1.CompanyList)     // 创建客户
		ApiRouter.POST("addCompanyList", v1.AddCompanyList)     // 增加客户
		ApiRouter.DELETE("companyList", v1.DeleteCompanylist)     // 删除客户
		ApiRouter.POST("updateCompanyList", v1.UpdateCompanylist)     // 删除客户
	}
}
