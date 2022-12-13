package view

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func AnyErrorRespond(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{Error: err.Error()})
}

func ErrorRespond(w http.ResponseWriter, err error) {
	AnyErrorRespond(w, http.StatusInternalServerError, err)
}
