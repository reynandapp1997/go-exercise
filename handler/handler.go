package handler

import "net/http"

func ResponseHandler(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	w.Write([]byte("{ \"message\": \"" + message + "\" }"))
}
