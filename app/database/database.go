package database

import "database/sql"

func NewDB() *sql.DB {

	sql.Open("mysql", "root:@tcp(localhost:3306)/")
}
