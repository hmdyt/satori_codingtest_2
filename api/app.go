package satori_codingtest_2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gorilla/mux"
	db "github.com/hmdyt/satori_codingtest-2/db"
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
	functions.HTTP("migrate", func(w http.ResponseWriter, r *http.Request) {
		client := db.GetDataBaseClient()
		defer client.Close()
		fmt.Println("start migration")


		if err := client.Schema.Create(context.Background()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("%#v /n", err.Error())
			json.NewEncoder(w).Encode(map[string]string{
				"msg": "migration failed",
			})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"msg": "migration success",
			})
		}
	})
}