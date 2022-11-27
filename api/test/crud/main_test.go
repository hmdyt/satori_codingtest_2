package main

import (
	"context"
	"os"
	"testing"

	"github.com/hmdyt/satori_codingtest-2/ent"
	"github.com/hmdyt/satori_codingtest-2/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {

	code := m.Run()

	os.Exit(code)
}

func ClearDB(client *ent.Client) {
	client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
}
