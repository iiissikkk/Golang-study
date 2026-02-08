package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	title string,
	description string,
	completed bool,
	createdAt time.Time,
) error {

	// Manual example (hardcoded values):
	// sqlQuery := `
	// INSERT INTO tasks (title, description, completed, created_at)
	// VALUES ('Breakfast', 'Morning Routine', FALSE, '2025-03-21 07:10:00')
	// `
	// _, err := conn.Exec(ctx, sqlQuery)
	// return err

	sqlQuery := `
	INSERT INTO tasks (
		title,
		description,
		completed,
		created_at
	)
	VALUES ($1, $2, $3, $4)
	`

	_, err := conn.Exec(ctx, sqlQuery, title, description, completed, createdAt)
	return err
}
