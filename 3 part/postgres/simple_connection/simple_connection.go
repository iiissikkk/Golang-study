package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	// Example Postgres connection:
	//ctx := context.Background()
	//
	//conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := conn.Ping(ctx); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Successfully connected to Database")

	return pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
}
