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
	router.HandleFunc("/user/{user_id}", handler.HandleGetUser).Methods("GET")
	router.HandleFunc("/user", handler.HandlePostUser).Methods("POST")
	router.HandleFunc("/mesuringpoint/{user_id}", handler.HandleGetMesuringPoints).Methods("GET")
	router.HandleFunc("/mesuringpoint", handler.HandlePostMesuringPoint).Methods("POST")
	router.HandleFunc("/mesuringpoint/{point_id}", handler.HandleDeleteMesuringPoint).Methods("DELETE")
	return router
}

func getPing(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "pong")
}
