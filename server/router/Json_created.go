package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCreatdJson(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("json").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("creatdJson", v1.CreatedJsonSql)     // 创建客户
		ApiRouter.GET("getTableHead", v1.GetTableHead)     // 获取表单头
	}
}
