package base

// FieldInfo 字段结构体，包含字段名称，是否为逐渐，字段对应变量值
type FieldInfo struct {
	FieldName  string      `json:"field_name"`  //字段名称
	PrimaryKey bool        `json:"primary_key"` //是否为主键
	ReValue    interface{} `json:"re_value"`    //字段对应变量值 通过反射获取，reflect.Value.Interface()
}

func (f *FieldInfo) InitFieldInfo(obj interface{}) (res []FieldInfo) {
	t, v := GetTypeOfAndValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i)
		if key.Tag.Get("PrimaryKey") == "true" {
			f.PrimaryKey = true
		} else {
			f.PrimaryKey = false
		}
		f.FieldName = key.Tag.Get("field")
		f.ReValue = v.FieldByName(key.Name).Interface()
		if f.checkNil() {
			res = append(res, *f)
		}

	}
	return
}

func (f *FieldInfo) checkNil() (res bool) {
	if f.ReValue == 0 || f.ReValue == "" {
		return
	}
	return true
}
