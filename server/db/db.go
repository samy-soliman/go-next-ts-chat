package db

import (
	"database/sql"
	"os"

	/*
		In Go, a blank import statement is used when you want to import a package solely for its side-effects, without using any of its exported identifiers. It’s written as _ "package/path".
		The underscore _ is known as the blank identifier. It’s a special identifier that you can use when you need to declare a variable that you won’t actually use.
		When you do a blank import, the Go runtime executes the imported package’s init function(s), if any.
		This can be useful when the init function sets up some state or performs some initialization that your program relies on.
		For example, in the case of _ "github.com/lib/pq", the pq package is a PostgreSQL driver for the database/sql package.
		The pq package’s init function registers the driver with the database/sql package.
		So, even though no identifiers from pq are directly used, it’s necessary to import pq for its side-effects.
		This allows you to use "postgres" as the driver name when calling sql.Open.
	*/
	_ "github.com/lib/pq"
)

type Database struct {
	// DB is a database handle representing a pool of zero or more underlying connections. It's safe for concurrent use by multiple goroutines.
	db *sql.DB //type DB struct
}

// a pointer is a variable that stores the memory address of another variable.
// Thats why we return &Database as its a memory value.
func NewDatabase() (*Database, error) {
	// func Open(driverName, dataSourceName string) (*DB, error)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
