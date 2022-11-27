package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hmdyt/satori_codingtest-2/crud"
	db "github.com/hmdyt/satori_codingtest-2/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateMesuringPoint(t *testing.T) {
	client := db.GetTestDatabaseClient()
	defer client.Close()
	ClearDB(client)

	point1, err := crud.CreateMesuringPoint(client, 1, 1.1)
	assert.Nil(t, err)
	assert.Equal(t, 1, point1.UserID)
	assert.Equal(t, 1.1, point1.BodyMass)

	point2, err := crud.CreateMesuringPoint(client, 1, 2.2)
	assert.Nil(t, err)
	assert.Equal(t, 1, point2.UserID)
	assert.Equal(t, 2.2, point2.BodyMass)

	point3, err := crud.CreateMesuringPoint(client, 2, 30)
	assert.Nil(t, err)
	assert.Equal(t, 2, point3.UserID)
	assert.Equal(t, 30., point3.BodyMass)

	points, err := crud.GetMesuringPoints(client, 1)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(points))
	assert.Equal(t, 1, points[0].ID)
	assert.Equal(t, 2, points[1].ID)
}

func TestDeleteMesuringPoint(t *testing.T) {
	client := db.GetTestDatabaseClient()
	defer client.Close()
	ClearDB(client)

	point1, err := crud.CreateMesuringPoint(client, 100, 1.1)
	assert.Nil(t, err)
	assert.Equal(t, 100, point1.UserID)
	assert.Equal(t, 1.1, point1.BodyMass)

	point1_get, err := crud.GetMesuringPoints(client, 100)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(point1_get))
	assert.Equal(t, 1, point1_get[0].ID)

	err = crud.DeleteMesuringPoint(client, 1)
	assert.Nil(t, err)

	point1_get, err = crud.GetMesuringPoints(client, 100)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(point1_get))
}
