package test

type Base struct {
	Id      int    `json:"id" PrimaryKey:"true" field:"id"` //主键
	Created string `json:"created" field:"created"`         //创建时间
	Updated string `json:"updated" field:"updated"`         //修改时间
}
