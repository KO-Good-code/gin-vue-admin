package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
)

// @Tags CreatedJsonSql
// @Summary 创建表单表头映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Table true "创建表单表头映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /json/creatdJson [post]
func CreatedJsonSql(c *gin.Context)  {
	var table model.Table
	_ = c.ShouldBindJSON(&table)
	if err := utils.Verify(table, utils.NameVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		
		return
	}
	if err := global.GVA_DB.Migrator().HasTable(table.Name); !err {
		service.CreateJsonTable(table.Name, table.Data)
	}
	if err := service.RemoveJsonTable(table.Name); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.UpdateJsonTable(table.Name, table.Data)
	response.OkWithMessage("操作成功", c);
}

// @Tags GetTableHead
// @Summary 获取表单表头映射
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Table true "预览创建代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /json/creatdJson [get]
func GetTableHead(c *gin.Context) {
	var table model.TableHead
	_ = c.ShouldBindQuery(&table)
	err, data := service.FindJsonTable(table.Name)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	scoreMap := make([]struct{
		Name string `json:"name"`
		Key string	`json:"key"`
	}, len(data))
	for i, key := range data {
		fmt.Println(i, key)
		scoreMap[i].Key = key.Key
		scoreMap[i].Name = key.Label
	}
	fmt.Println(scoreMap)
	response.OkWithData(scoreMap, c)
}