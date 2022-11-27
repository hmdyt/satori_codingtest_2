package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hmdyt/satori_codingtest-2/crud"
	db "github.com/hmdyt/satori_codingtest-2/db"
	"github.com/hmdyt/satori_codingtest-2/model"
)

func HandlePostUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userPostRequest model.UserPostRequest
	err := decodeRequest(request, &userPostRequest)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	client := db.GetDataBaseClient()
	defer client.Close()
	user, err := crud.CreateUser(client, userPostRequest.Name, userPostRequest.Email)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func HandleGetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)

	path_params := mux.Vars(request)
	user_id, err := strconv.Atoi(path_params["user_id"])
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
	}

	client := db.GetDataBaseClient()
	defer client.Close()
	user := crud.GetUser(client, user_id)
	json.NewEncoder(writer).Encode(user)
}
