package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		unexpectedErr(w, errors.New("JSON marshal failed"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}

// err write error, status to http response writer
func unexpectedErr(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	http.Error(w, err.Error(), status)
}
