package common

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func SendResponse(data any, w http.ResponseWriter, r *http.Request) {
	jsonResponse, err := json.Marshal(JsonResponse{Success: true, Data: data, Message: ""})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Установка HTTP-заголовка Content-Type на application/json
	w.Header().Set("Content-Type", "application/json")

	// Запись JSON-ответа в http.ResponseWriter
	w.Write(jsonResponse)
}
