package operation

import "database/sql"

type Service interface {
	Insert() (sql.Result, error)
	QueryList() (*sql.Rows, error)
	Update() (sql.Result, error)
	Delete() (sql.Result, error)
	Query() (*sql.Rows, error)
	Count() (int, error)
}
