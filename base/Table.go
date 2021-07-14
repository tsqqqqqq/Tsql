package base

import (
	"fmt"
	"reflect"
)

// TableInfo 表操作结构体
type TableInfo struct {
	TableName string      `json:"table_name"` //表名
	Fields    []FieldInfo `json:"fields"`     //表中的字段切片
}

func (t *TableInfo) InitTableInfo(obj interface{}) (res *TableInfo) {
	ty, _ := GetTypeOfAndValueOf(obj)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	t.TableName = ty.Name()
	f := new(FieldInfo)
	t.Fields = f.InitFieldInfo(obj)
	fmt.Println(t)
	return t
}
