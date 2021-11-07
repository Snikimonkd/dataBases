package models

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/Snikimonkd/dataBases/internal/utils/metrics"

	"net/http"
)

func MetricFunc(code int, r *http.Request, err error) {
	if code >= 500 {
		metrics.CreateRequestErrors(r, err)
		return
	}
	metrics.CreateRequestHits(code, r)
}

func ResponseJson(res interface{}, status int, w http.ResponseWriter, r *http.Request) {
	MetricFunc(status, r, nil)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

func ResponseError(str string, status int, w http.ResponseWriter, r *http.Request) {
	MetricFunc(status, r, errors.New(str))
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
