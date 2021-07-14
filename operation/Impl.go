package operation

import (
	"TSQL/base"
	"database/sql"
	"fmt"
	"strings"
)

const (
	where = " where "
	set   = " set "
	order = " order by "
	group = " group by "
	and   = " and "
	or    = " or "
	on    = " on "
)

type Impl struct {
	tables    *base.TableInfo
	value     []interface{}
	filedSize int
	condition string
}

func (i *Impl) Init(obj interface{}) {
	i.tables = new(base.TableInfo).InitTableInfo(obj)
	i.filedSize = len(i.tables.Fields)
}

func (i *Impl) Insert() (res sql.Result) {
	fields := i.tables.Fields
	len := i.filedSize
	var value string
	sql := "insert into " + i.tables.TableName + "("
	for j := 0; j < len; j++ {
		sql += fields[j].FieldName + ","
		value += "?,"
		i.value = append(i.value, fields[j].ReValue)
	}
	sql = strings.TrimRight(sql, ",")
	sql += ")  value(" + strings.TrimRight(value, ",") + ")"
	fmt.Println(sql, i.value)
	return res
}

func (i *Impl) Delete() (res sql.Result) {
	sql := "delete from "
	fmt.Println(sql)
	return
}

func (i *Impl) QueryById(id int) (res sql.Result) {
	return
}

func (i *Impl) Update() (res sql.Result) {
	fields := i.tables.Fields
	len := i.filedSize
	sql := "update " + i.tables.TableName + " set "
	for j := 0; j < len; j++ {
		sql += fields[j].FieldName + "=?, "
		i.value = append(i.value, fields[j].ReValue)
	}
	sql = strings.TrimRight(sql, ", ")
	sql += i.condition
	fmt.Println(sql, i.value)
	return res
}

func (i *Impl) QueryList() (res sql.Result) {
	return
}

func (i *Impl) Count() (res sql.Result) {
	return
}

func (i *Impl) Where(field string, value interface{}) (res *Impl) {
	if i.checkCondition() {
		i.condition = where + field + "= " + value.(string)
	}
	return i
}

func (i *Impl) And() (res *Impl) {

	return i
}

func (i *Impl) setValue() {

}

func (i *Impl) checkCondition() bool {
	if i.condition == "" || strings.EqualFold(i.condition, "") {
		return true
	}
	return false
}
