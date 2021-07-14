package main

import (
	"TSQL/operation"
	"TSQL/test"
)

func main() {
	stu := test.Student{
		Name:     "滕盛强",
		Age:      23,
		Sex:      1,
		Birthday: "1998-03-25",
		ClassId:  2002,
	}
	Service := new(operation.Impl)
	Service.Init(stu)
	Service.Where("id", "1").Update()
	//f :=new(base.TableInfo)
	//f.InitTableInfo(stu)
}
