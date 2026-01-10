package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statuscode int) {
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
