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