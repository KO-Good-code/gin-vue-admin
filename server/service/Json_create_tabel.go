package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//@author: [chenrui](https://github.com/piexlmax)
//@function: CreateJsonTable
//@description: 创建表单对应数据表
//@param: name string, data model.TableMap
//@return: error
func CreateJsonTable(name string, data []model.TableMap) error {
	err := global.GVA_DB.Table(name).Migrator().CreateTable(&data)
	return err
}

//@author: [chenrui](https://github.com/piexlmax)
//@function: UpdateJsonTable
//@description: 修改表单对应数据表
//@param: name string, data model.TableMap
//@return: error
func UpdateJsonTable(name string, data []model.TableMap) error {
	err := global.GVA_DB.Table(name).Create(&data).Error
	return err
}

//@author: [chenrui](https://github.com/piexlmax)
//@function: RemoveJsonTable
//@description: 删除表单对应数据表
//@param: name string, data model.TableMap
//@return: error
func RemoveJsonTable(name string) error {
	var table model.TableMap
 	err := global.GVA_DB.Table(name).Where("1=1").Delete(&table).Error
	return err
}

//@author: [chenrui](https://github.com/piexlmax)
//@function: FindJsonTable
//@description: 查询表单对应数据表
//@param: name string
//@return: error
func FindJsonTable(name string) (err error, json []model.TableMap) {
	var list []model.TableMap
	err = global.GVA_DB.Table(name).Find(&list).Error
	return err, list
}