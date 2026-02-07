package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
