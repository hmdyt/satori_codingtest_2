package satori_codingtest_2

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmdyt/satori_codingtest-2/handler"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", getPing).Methods("GET")
	router.HandleFunc("/user", handler.HandleGetUser).Methods("GET")
	return router
}

func getPing(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "pong")
}
