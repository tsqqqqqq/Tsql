package operation

import "database/sql"

type Service interface {
	Insert() sql.Result
	QueryList() sql.Result
	Update() sql.Result
	Delete(id int) sql.Result
	QueryById(id int) sql.Result
	Count() sql.Result
}
