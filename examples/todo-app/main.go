package todoapp

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/conneroisu/dbpu"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	mtx         *sync.Mutex
	shuffle     *sync.Once
}

func pingContext(ctx context.Context, db *sql.DB) {
	err := db.PingContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ping db: %s", err)
		os.Exit(1)
	}
}

func exec(ctx context.Context, db *sql.DB, stmt string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, stmt, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", stmt, err)
		os.Exit(1)
	}
	return res
}
func execTx(ctx context.Context, tx *sql.Tx, stmt string, args ...any) sql.Result {
	res, err := tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", stmt, err)
		os.Exit(1)
	}
	return res
}

func queryTx(ctx context.Context, tx *sql.Tx, stmt string, args ...any) *sql.Rows {
	res, err := tx.QueryContext(ctx, stmt, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query %s: %s", stmt, err)
		os.Exit(1)
	}
	return res
}

func main() {
	// Create a new database

	db, err := dbpu.CreateDatabase("orgToken", "orgName", "name", "group")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create a new token for the database
	token, err := dbpu.CreateDatabaseToken("orgName", "dbName", "apiTok")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the token
	fmt.Println(token)

	// print the database
	fmt.Println(db)

	// connect to the database
	condb, err := sql.Open("libsql", db.Hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", db.Hostname, err)
		os.Exit(1)
	}
	ctx := context.Background()

	// print the database connection
	fmt.Println(condb)
	// print the context
	fmt.Println(ctx)

	// ping the database
	pingContext(ctx, condb)

}
