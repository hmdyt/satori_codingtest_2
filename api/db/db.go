package satori_codingtest_2

import (
	"fmt"
	"os"

	"github.com/hmdyt/satori_codingtest-2/ent"
)

func GetDataBaseClient() *ent.Client {
	client, err := ent.Open("mysql", "testUser:testPassword@tcp(mysql:3306)/testDatabase?parseTime=True")
	if err != nil {
		fmt.Printf("DB Open Error: %s \n", err.Error())
		os.Exit(1)
	}
	return client
}

func GetTestDatabaseClient() *ent.Client {
	os.Remove("test.sqlite")
	client, err := ent.Open("sqlite3", "file:test.sqlite?mode=memory&cache=shared&_fk=1")
	if err != nil {
		fmt.Printf("DB Open Error: %s \n", err.Error())
		os.Exit(1)
	}
	return client
}
