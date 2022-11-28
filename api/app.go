package satori_codingtest_2

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = Router()
	functions.HTTP("hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Request-Method", "*")
		w.Header().Set("Content-Type", "*")
		if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
		router.ServeHTTP(w, r)
	})
}