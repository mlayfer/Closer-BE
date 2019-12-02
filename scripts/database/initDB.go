package main

import (
	"Closer/dataaccess"
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, _ = db.Exec("DROP SCHEMA CloserDB")
	_, _ = db.Exec("CREATE SCHEMA CloserDB")

	closerDB, closer := dataaccess.NewDatabaseAccess("root", "root")
	defer closer()

	closerDB.InitDB()
}