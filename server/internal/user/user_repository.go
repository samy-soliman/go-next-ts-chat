package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	//type sql.DB is a struct that implements all methods in DBTX interface so it implements the interface implicitly and can be bassed here as argument.
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

/*
In your NewRepository function, the return type is Repository, which is an interface.
An interface in Go can be satisfied by any type that implements the methods declared in the interface, whether it’s a value type or a pointer type.

The & operator is used to create a pointer to the repository struct.
This means &repository{db: db} is of type *repository (a pointer to a repository).

Now, if your *repository type implements all the methods required by the Repository interface, then a *repository can be returned where a Repository is expected.
This is why you’re able to return &repository{db: db} even though the function signature indicates it returns a Repository.
*/
/*
In Go, the methods that a type needs to implement an interface are determined by the receiver type of the methods.
If the methods are defined with a pointer receiver (like *repository), then a pointer to the type (not the value of the type itself) is needed to satisfy the interface.

In your case, the Repository interface requires methods like CreateUser and GetUserByEmail, which are defined with a pointer receiver (*repository).
This means that these methods are associated with *repository (pointer to repository), not repository itself.

So, when you try to return a repository value from NewRepository, it doesn’t satisfy the Repository interface because the required methods are not associated with repository. They’re associated with *repository.

That’s why you need to return a pointer to repository (i.e., &repository{db: db}) from NewRepository.
The & operator gets the memory address of the repository value (i.e., it creates a pointer to the repository), and this pointer does satisfy the Repository interface.
*/
func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	// func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId) //Rows.Scan
	// func (r *Row) Scan(dest ...any) error >>> Scan copies the columns from the matched row into the values pointed at by dest.
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}
