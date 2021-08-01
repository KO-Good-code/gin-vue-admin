package v1

import (
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
)

func CreatedJsonSql(c *gin.Context)  {
	var a string = "agshdgahj"
	response.OkWithData(a, c);
}
