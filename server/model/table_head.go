package model

import (
	// "gin-vue-admin/global"
)

// 
type TableHead struct {
	Name string `json:"name" form:"name"`
}

type TableHeadMap struct {
	Name string
	Key string
}