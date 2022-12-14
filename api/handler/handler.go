package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"

	_ "github.com/go-sql-driver/mysql"
)

var decoder = schema.NewDecoder()

func decodeRequest(request *http.Request, dst interface{}) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}

	err = decoder.Decode(dst, request.PostForm)
	if err != nil {
		return err
	}

	fmt.Printf("yet parsed: %#v \n", request.PostForm)
	fmt.Printf("parsed: %#v \n", dst)


	return nil
}
