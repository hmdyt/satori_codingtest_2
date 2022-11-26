package satori_codingtest_2

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = Router()
	functions.HTTP("hello", router.ServeHTTP)
}

// func entryPoint(writer http.ResponseWriter, request *http.Request) {
// 	router.ServeHTTP()
// }
