package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hmdyt/satori_codingtest-2/crud"
	db "github.com/hmdyt/satori_codingtest-2/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	client := db.GetTestDatabaseClient()
	defer client.Close()
	ClearDB(client)

	user1, err := crud.CreateUser(client, "name1", "email1")
	assert.Nil(t, err)
	assert.Equal(t, 1, user1.ID)
	assert.Equal(t, "name1", user1.Name)
	assert.Equal(t, "email1", user1.Email)

	user2, err := crud.CreateUser(client, "name2", "email2")
	assert.Nil(t, err)
	assert.Equal(t, 2, user2.ID)
	assert.Equal(t, "name2", user2.Name)
	assert.Equal(t, "email2", user2.Email)

	user1_get := crud.GetUser(client, 1)
	assert.Equal(t, 1, user1_get.ID)
	assert.Equal(t, "name1", user1_get.Name)
	assert.Equal(t, "email1", user1_get.Email)

	user2_get := crud.GetUser(client, 2)
	assert.Equal(t, user2_get.ID, 2)
	assert.Equal(t, "name2", user2_get.Name)
	assert.Equal(t, "email2", user2_get.Email)
}
