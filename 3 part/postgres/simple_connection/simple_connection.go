package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func CheckConnection() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Database")
}
