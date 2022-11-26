package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hmdyt/satori_codingtest-2/model"
)

func HandleGetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(model.User{Id: 1, Name: "Aaa Bbb"})
}
