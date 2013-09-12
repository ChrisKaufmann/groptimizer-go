package main

import (
	"database/sql"
)

var (
	stmtGetTrip  *sql.Stmt
	stmtSaveTrip *sql.Stmt
	stmtFindTrip *sql.Stmt
	stmtAddTrip  *sql.Stmt
)

func init() {
	stmtGetTrip = sth(db, "select id,date,store_id from trip where id=?")
	stmtSaveTrip = sth(db, "update trip set store_id=? where id=?")
	stmtFindTrip = sth(db, "select id,date,store_id from trip where  like '%?%'")
	stmtAddTrip = sth(db, "insert into trip (name) values (?)")
}

type Trip struct {
	ID   int
	Date string
	Store Store
}

func (i Trip) Insert() {
	stmtAddTrip.Exec(i.Store.ID)
}

func (i Trip) Save() {
	stmtSaveTrip.Exec(i.Store.ID)
}
func getTrip(id string) (t Trip, err error) {
	var sid string
	err = stmtGetTrip.QueryRow(id).Scan(&t.ID, &t.Date,&sid)
	if err != nil {
		err.Error()
	}
	s,err := getStore(sid)
	if err != nil {
		var tt Trip
		return tt, err
	}
	t.Store=s
	return t, err
}
