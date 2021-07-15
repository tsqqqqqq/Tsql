package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var err error

type DbConfig struct {
	Db  *sql.DB
	url string
}

func (db *DbConfig) Connect(config interface{}) {
	db.build(config)
	db.Db, err = sql.Open("mysql", db.url)
	if err != nil {
		panic(err)
	}
}

func (db *DbConfig) build(config interface{}) {
	_, v := GetTypeOfAndValueOf(config)
	ip := v.FieldByName("Ip").Interface()
	port := v.FieldByName("Port").Interface()
	username := v.FieldByName("Username").Interface()
	password := v.FieldByName("Password").Interface()
	databaseName := v.FieldByName("DatabaseName").Interface()
	db.url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, ip, port, databaseName)
}

func (db *DbConfig) Close() {
	db.Db.Close()
}

func (db *DbConfig) Query(sql string, obj []interface{}) (*sql.Rows, error) {
	rows, err := db.Db.Query(sql, obj...)
	return rows, err
}

func (db *DbConfig) ExecData(sql string, obj []interface{}) (res sql.Result, err error) {
	fmt.Println("Exec", sql, cap(obj))
	res, err = db.Db.Exec(sql, obj...)
	return
}
