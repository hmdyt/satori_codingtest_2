package satori_codingtest_2

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/hmdyt/satori_codingtest-2/handler"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/{path:.*}", handleProxy(getProxyNextJs())).Methods("GET")
	router.HandleFunc("/ping", getPing).Methods("GET")
	router.HandleFunc("/user/{user_id}", handler.HandleGetUser).Methods("GET")
	router.HandleFunc("/user", handler.HandlePostUser).Methods("POST")
	router.HandleFunc("/mesuringpoint/{user_id}", handler.HandleGetMesuringPoints).Methods("GET")
	router.HandleFunc("/mesuringpoint", handler.HandlePostMesuringPoint).Methods("POST")
	router.HandleFunc("/mesuringpoint/{point_id}", handler.HandleDeleteMesuringPoint).Methods("DELETE")
	return router
}

func getPing(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(map[string]string{"ping": "pong"})
}

func getProxyNextJs() *httputil.ReverseProxy{
	target_url := os.Getenv("NEXT_HOST")
	u, err := url.Parse(target_url)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	return proxy
}

func handleProxy(proxy *httputil.ReverseProxy) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		request.URL.Path = mux.Vars(request)["path"]
		proxy.ServeHTTP(writer, request)
	}
}