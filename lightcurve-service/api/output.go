package api

import (
	"encoding/json"
	"net/http"
)

func FormatOutput[T any](result T, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(result)
}
