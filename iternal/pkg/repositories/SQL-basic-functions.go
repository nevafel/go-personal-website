/*
Some SQL tools to reduce boilerplate.
Probably will be rewritten in future.
*/
package repositories

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dixxe/personal-website/iternal/pkg"
)

func OpenDb(database_name string) *sql.DB {
	if _, err := os.Stat("./" + database_name); err != nil {
		panic(database_name + " not found!")
	}
	db, err := sql.Open("sqlite", database_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database was openned")

	return db
}

// I don't know a way how to automate this process.
func InitDb(repo pkg.Repository[Post]) {
	//defer db.Close()

	repo.ExecSpecific(`
	CREATE TABLE blogs(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		header TEXT,
		content TEXT
	);
	`)
	fmt.Println("Database was initiated.")
}
