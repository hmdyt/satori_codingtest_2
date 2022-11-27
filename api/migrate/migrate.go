package main

import (
	"context"
	"log"

	db "github.com/hmdyt/satori_codingtest-2/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client := db.GetDataBaseClient()
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
