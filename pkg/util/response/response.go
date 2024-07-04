package response

import (
	"encoding/json"
	"errors"
	"net/http"

	error2 "nam_0508/pkg/error"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		unexpectedErr(w, errors.New("JSON marshal failed"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}

func Error(w http.ResponseWriter, err interface{}) {
	switch err.(type) {
	case *error2.XError:
		w.WriteHeader(err.(error2.XError).GetHTTPStatus())
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	b, err := json.Marshal(err)
	if err != nil {
		unexpectedErr(w, errors.New("JSON marshal failed"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, _ = w.Write(b)
}

// err write error, status to http response writer
func unexpectedErr(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	http.Error(w, err.Error(), status)
}
