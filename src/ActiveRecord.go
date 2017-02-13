package src

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ActiveRecord struct {
	isNew bool
	table string
}

func (ar *ActiveRecord) SetTable(table string) {
	ar.table = table
}

func (ar *ActiveRecord) SetIsNewRecord(is bool) {
	ar.isNew = is
}

func (ar *ActiveRecord) FindOne(id string) *sql.Row {
	db, err := OpenDB()
	CheckErr(err)
	defer db.Close()
	return db.QueryRow("SELECT * FROM " + ar.table + " WHERE id=" + id)
}

func (ar *ActiveRecord) Find(where string) (*sql.Rows, error) {
	db, err := OpenDB()
	CheckErr(err)
	defer db.Close()
	if len(where) > 0 {
		return db.Query("SELECT * FROM " + ar.table + " WHERE " + where)
	} else {
		return db.Query("SELECT * FROM " + ar.table)
	}
}
