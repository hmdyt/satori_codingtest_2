package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/hmdyt/satori_codingtest-2/crud"
	db "github.com/hmdyt/satori_codingtest-2/db"
	"github.com/hmdyt/satori_codingtest-2/model"
)

func HandlePostMesuringPoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var mesuringPointPostRequest model.MesuringPointPostRequest
	err := decodeRequest(request, &mesuringPointPostRequest)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	client := db.GetDataBaseClient()
	defer client.Close()
	point, err := crud.CreateMesuringPoint(
		client,
		mesuringPointPostRequest.User_id,
		mesuringPointPostRequest.Body_mass,
	)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewEncoder(writer).Encode(point)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func HandleGetMesuringPoints(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	path_params := mux.Vars(request)
	user_id, err := strconv.Atoi(path_params["user_id"])
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
	}

	client := db.GetDataBaseClient()
	defer client.Close()
	points, err := crud.GetMesuringPoints(client, user_id)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
	}

	json.NewEncoder(writer).Encode(points)
	writer.WriteHeader(http.StatusAccepted)
}

func HandleDeleteMesuringPoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	path_params := mux.Vars(request)
	point_id, err := strconv.Atoi(path_params["point_id"])
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
	}

	client := db.GetDataBaseClient()
	defer client.Close()
	err = crud.DeleteMesuringPoint(client, point_id)
	if err != nil {
		model.WriteError(writer, http.StatusBadRequest, err.Error())
		return
	}

	writer.WriteHeader(http.StatusAccepted)
}
