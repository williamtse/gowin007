package src

import (
	"database/sql"
)

type League struct {
	id    int
	name  string
	color string
	ar    ActiveRecord
}

func (league *League) Save() (sql.Result, error) {
	db, err := OpenDB()
	CheckErr(err)
	defer db.Close()
	if !league.ar.isNew {
		stmt, err := db.Prepare("INSERT INTO " +
			league.ar.table + " (id,name,color)values(?,?,?)")
		CheckErr(err)
		res, err := stmt.Exec("null", league.name, league.color)
		return res, err
	} else {
		stmt, err := db.Prepare("UPDATE " + league.ar.table + " set name=?,color=? WHERE id=" + string(league.id))
		CheckErr(err)
		res, err := stmt.Exec(league.name, league.color)
		return res, err
	}

}
