package main

import (
	"database/sql"
)

var (
	stmtFindStore *sql.Stmt
	stmtGetStore *sql.Stmt
	stmtSaveStore *sql.Stmt
	stmtAddStore *sql.Stmt
)

func init() {
	stmtFindStore=sth(db,"select id,name from store where name like '%?%'")
	stmtGetStore=sth(db,"select id,name from store where id=?")
	stmtSaveStore=sth(db,"update store set name=? where id=?")
	stmtAddStore=sth(db,"insert into store (name) values (?)")
}

func (s Store)Save () {
	stmtSaveStore.Exec(s.Name,s.ID)
}
type Store struct{
	ID int
	Name string
}
func (s Store) Insert() {
	if s.Name=="" {
		panic("Name is blank for new store")
	}
	stmtAddStore.Exec(s.Name)
}
func getStore(id string) (s Store, err error) {
	err = stmtGetStore.QueryRow(id).Scan(&s.ID,&s.Name)
	if err != nil {
		err.Error()
	}
	return s,err
}
func findStore(n string) []Store {
	var sl []Store
	rows,_ := stmtFindStore.Query(n)

	for rows.Next() {
		var s Store
		rows.Scan(&s.ID,&s.Name)
		sl = append(sl, s)
	}
	return sl
}
