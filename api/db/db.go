package satori_codingtest_2

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/hmdyt/satori_codingtest-2/ent"
)

func GetDataBaseClient() *ent.Client {
	mysql_config := mysql.Config{
		DBName: os.Getenv("MYSQL_DATABASE"),
		User: os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Addr: os.Getenv("MYSQL_HOST"),
		Net: os.Getenv("MYSQL_NET"),
		ParseTime: true,
		AllowNativePasswords: true,
	}
	
	client, err := ent.Open("mysql", mysql_config.FormatDSN())
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
