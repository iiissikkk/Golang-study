package main

import (
	"context"
	"fmt"
	"postgresPractice/simple_connection"
	"postgresPractice/simple_sql"
	"time"

	"github.com/k0kubun/pp/v3"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	//if err := simple_sql.InsertRow(
	//	ctx,
	//	conn,
	//	"Dinner",
	//	"Shwarma",
	//	false,
	//	time.Now(),
	//); err != nil {
	//	panic(err)
	//}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	pp.Println(tasks)

	for _, task := range tasks {
		if task.ID == 6 {
			task.Title = "Lunch"
			task.Description = "Egg"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}

	//if err := simple_sql.DeleteRow(ctx, conn); err != nil {
	//	panic(err)
	//}

	tasks1, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	pp.Println(tasks1)

	fmt.Println("Succeed")
}
