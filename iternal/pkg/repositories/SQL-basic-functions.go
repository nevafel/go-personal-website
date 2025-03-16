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

type blankBlog struct {
	Database *sql.DB
}

func (blank blankBlog) GetAllValues() ([]Post, error)         { return nil, nil }
func (blank blankBlog) GetValueByID(id int) (Post, error)     { return Post{}, nil }
func (blank blankBlog) InsertValue(Post) (id int, err error)  { return 0, nil }
func (blank blankBlog) DeleteValueByID(id int) error          { return nil }
func (blank blankBlog) UpdateValue(id int, object Post) error { return nil }
func (blank blankBlog) ExecSpecific(SQL_command string) (sql.Result, error) {
	result, err := blank.Database.Exec(SQL_command)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (blank blankBlog) QuerySpecific(SQL_command string) (*sql.Rows, error) { return nil, nil }

func OpenDb(database_name string) *sql.DB {
	if _, err := os.Stat(database_name); err != nil {
		fmt.Println(database_name + " not found! Creating a blank one...")

		if _, err := os.Create(database_name); err != nil {
			panic("Failed to create a database!")
		}
		blank := blankBlog{Database: OpenDb(database_name)}
		initDb(blank)

		return blank.Database
	}
	db, err := sql.Open("sqlite", database_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database was openned")

	return db
}

// I don't know a way how to automate this process.
// UPDATE: figured it out. Now blog creates automaticly if it can't locate it
func initDb(repo pkg.Repository[Post]) {
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

func repairDb(repo pkg.Repository[Post], columnName string, columnType string, defaultValue string) {

	command := fmt.Sprintf("SELECT %v FROM blogs;", columnName)

	if _, err := repo.QuerySpecific(command); err == nil {
		fmt.Printf("Database has column %v. Repairing not needed", columnName)
		return
	}

	// Currently repo.ExecSpecific can't provide security against SQL injections.
	command = fmt.Sprintf("ALTER TABLE blogs ADD %v %v;", columnName, columnType)
	repo.ExecSpecific(command)

	command = fmt.Sprintf("UPDATE blogs SET %v = %v;", columnName, defaultValue)
	repo.ExecSpecific(command)

	fmt.Printf("Database blogs was modified to be unified. Added %v with type %v\n",
		columnName, columnType)

}
