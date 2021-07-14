package base

import "reflect"

// FieldInfo 字段结构体，包含字段名称，是否为逐渐，字段对应变量值
type FieldInfo struct {
	FieldName  string        //字段名称
	PrimaryKey bool          //是否为主键
	ReValue    reflect.Value //字段对应变量值 通过反射获取，reflect.Value.Interface()
}

func test() string {
	return "abc"
}
