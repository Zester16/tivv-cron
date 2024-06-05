package utils

import (
	"encoding/json"
	"net/http"
)

type errorMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func SetErrorForNoQuery(w http.ResponseWriter) {
	sendResponse400(w, errorMessage{Message: "no query passed", StatusCode: 1})
}

func SetErrorForNoMatchingQuery(w http.ResponseWriter) {
	sendResponse400(w, errorMessage{Message: "Wrong find", StatusCode: 2})
}

func SetError400(w http.ResponseWriter, err error) {
	sendResponse400(w, errorMessage{Message: err.Error(), StatusCode: 3})
}

func sendResponse400(w http.ResponseWriter, em errorMessage) {
	str, _ := json.Marshal(em)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(str)
}
