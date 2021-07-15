package operation

import (
	"database/sql"
	"fmt"
	"github.com/tsqqqqqq/Tsql/base"
	"strings"
)

var value string

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
	tables     *base.TableInfo
	value      []interface{}
	filedSize  int
	condition  string
	db         *base.DbConfig
	whereState int
}

func (i *Impl) Init(obj interface{}) {
	i.tables = new(base.TableInfo).InitTableInfo(obj)
	i.filedSize = len(i.tables.Fields)
}

func (i *Impl) DbInit(dbconfig interface{}) {
	i.db = new(base.DbConfig)
	i.db.Connect(dbconfig)
}

func (i *Impl) Insert() (res sql.Result, err error) {
	fields := i.tables.Fields
	len := i.filedSize
	sql := "insert into " + i.tables.TableName + "("
	for j := 0; j < len; j++ {
		sql += fields[j].FieldName + ","
		value += "?,"
		i.value = append(i.value, fields[j].ReValue)
	}
	sql = strings.TrimRight(sql, ",")
	sql += ")  value(" + strings.TrimRight(value, ",") + ")"
	res, err = i.db.ExecData(sql, i.value)
	if err != nil {
		panic(err)
	}
	fmt.Println(sql, i.value)
	return res, err
}

func (i *Impl) Delete() (res sql.Result, err error) {
	sql := "delete from " + i.tables.TableName + i.condition
	fmt.Println(sql)
	res, err = i.db.ExecData(sql, i.value)
	return
}

func (i *Impl) Query() (res *sql.Rows, err error) {
	sql := "select * from " + i.tables.TableName + " " + i.condition
	fmt.Println(sql)
	res, err = i.db.Query(sql, i.value)
	return
}

func (i *Impl) Update() (res sql.Result, err error) {
	fields := i.tables.Fields
	len := i.filedSize
	sql := "update " + i.tables.TableName + set
	for j := 0; j < len; j++ {
		sql += fields[j].FieldName + "=?, "
		i.value = append(i.value, fields[j].ReValue)
	}
	sql = strings.TrimRight(sql, ", ")
	sql += i.condition
	res, err = i.db.ExecData(sql, i.value)
	fmt.Println(sql, i.value)
	return res, err
}

func (i *Impl) QueryList() (res *sql.Rows, err error) {
	return
}

func (i *Impl) Count() (res int, err error) {
	var primaryKeyName string
	for j := 0; j < len(i.tables.Fields); j++ {
		if i.tables.Fields[j].PrimaryKey {
			primaryKeyName = i.tables.Fields[j].FieldName
		}
	}
	sql := "select count(" + primaryKeyName + ") from " + i.tables.TableName
	fmt.Println(sql)
	rows, err := i.db.Query(sql, nil)
	var count int
	if rows.Next() {
		err = rows.Scan(&count)
	}
	return count, err
}

func (i *Impl) Where(field string, value interface{}) (res *Impl) {
	fmt.Println(i.whereState)
	if i.checkCondition() {
		i.condition = where + field + " = ?" + " "
		i.whereState = 1
		i.value = append(i.value, value)
		return i
	}
	if i.whereState != 0 {
		i = i.SetFieldAndValue(field, value)
	}
	return i
}

func (i *Impl) And() (res *Impl) {
	i.condition += and + " "
	return i
}

func (i *Impl) Or() (res *Impl) {
	i.condition += or + " "
	return i
}

func (i *Impl) SetFieldAndValue(field string, value interface{}) (res *Impl) {
	i.condition += field + " = ? "
	i.value = append(i.value, value)
	return i
}

func (i *Impl) checkCondition() bool {
	if i.condition == "" || strings.EqualFold(i.condition, "") {
		return true
	}
	return false
}
