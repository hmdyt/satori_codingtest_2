package model

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int
	Reason     string
}

func WriteError(writer http.ResponseWriter, httpStatus int, reason string) {
	writer.WriteHeader(httpStatus)
	jsonResponse, _ := json.Marshal(ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Reason:     reason,
	})
	writer.Write(jsonResponse)
}
