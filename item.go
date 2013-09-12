package main

import (
	"database/sql"
)

var (
	stmtGetItem  *sql.Stmt
	stmtSaveItem *sql.Stmt
	stmtFindItem *sql.Stmt
	stmtAddItem  *sql.Stmt
)

func init() {
	stmtGetItem = sth(db, "select id,name from item where id=?")
	stmtSaveItem = sth(db, "update item set name=? where id=?")
	stmtFindItem = sth(db, "select id,name from item where name like '%?%'")
	stmtAddItem = sth(db, "insert into item (name) values (?)")
}

type Item struct {
	ID   int
	Name string
}

func (i Item) Insert() {
	if i.Name == "" {
		panic("Name is blank for new item")
	}
	stmtAddItem.Exec(i.Name)
}

func (i Item) Save() {
	stmtSaveItem.Exec(i.Name, i.ID)
}
func getItem(id string) (i Item, err error) {
	err = stmtGetItem.QueryRow(id).Scan(&i.ID, &i.Name)
	if err != nil {
		err.Error()
	}
	return i, err
}
func findItem(n string) []Item {
	var il []Item
	rows, _ := stmtFindItem.Query(n)

	for rows.Next() {
		var i Item
		rows.Scan(&i.ID, &i.Name)
		il = append(il, i)
	}
	return il
}
