package models

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJson(res interface{}, status int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

func ResponseError(str string, status int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	res := ModelError{
		Message: str,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}
