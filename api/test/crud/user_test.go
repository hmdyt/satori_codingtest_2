package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hmdyt/satori_codingtest-2/crud"
	db "github.com/hmdyt/satori_codingtest-2/db"
	"github.com/hmdyt/satori_codingtest-2/ent"
	"github.com/stretchr/testify/assert"
)

var client *ent.Client

func TestMain(m *testing.M) {
	client = db.GetTestDatabaseClient()

	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Printf("failed creating schema resources: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	user1, err := crud.CreateUser(client, "name1", "email1")
	assert.Nil(t, err)
	assert.Equal(t, user1.ID, 1)
	assert.Equal(t, user1.Name, "name1")
	assert.Equal(t, user1.Email, "email1")

	user2, err := crud.CreateUser(client, "name2", "email2")
	assert.Nil(t, err)
	assert.Equal(t, user2.ID, 2)
	assert.Equal(t, user2.Name, "name2")
	assert.Equal(t, user2.Email, "email2")

	user1_get := crud.GetUser(client, 1)
	assert.Equal(t, user1_get.ID, 1)
	assert.Equal(t, user1_get.Name, "name1")
	assert.Equal(t, user1_get.Email, "email1")

	user2_get := crud.GetUser(client, 2)
	assert.Equal(t, user2_get.ID, 2)
	assert.Equal(t, user2_get.Name, "name2")
	assert.Equal(t, user2_get.Email, "email2")
}
