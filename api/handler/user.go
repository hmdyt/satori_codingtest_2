package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hmdyt/satori_codingtest-2/crud"
	db "github.com/hmdyt/satori_codingtest-2/db"
	"github.com/hmdyt/satori_codingtest-2/model"
)

var decoder = schema.NewDecoder()

func HandlePostUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	err := request.ParseForm()
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	var userPostRequest model.UserBase
	err = decoder.Decode(&userPostRequest, request.PostForm)
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
	user_id, err := strconv.Atoi(path_params["id"])
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
	}

	client := db.GetDataBaseClient()
	defer client.Close()
	user := crud.GetUser(client, user_id)
	json.NewEncoder(writer).Encode(user)
}
