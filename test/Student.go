package test

type Student struct {
	Id       int    `json:"id" PrimaryKey:"true" field:"id"` //主键
	Name     string `json:"age" field:"name"`
	Age      int    `json:"age" field:"age"`
	Sex      int    `json:"sex" field:"sex"`
	Birthday string `json:"birthday" field:"birthday"`
	ClassId  int    `json:"class_id" field:"class_id"`
	Created  string `json:"created" field:"created"` //创建时间
	Updated  string `json:"updated" field:"updated"` //修改时间
}
