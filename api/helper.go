package api

import (
	"net/http"
	"encoding/json"
)

func WriteJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json encoding=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
