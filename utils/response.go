package utils

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

func ResponseErr(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte("{\"Error\": \"" + message + "\"}"))
}

func ResponseDone(w http.ResponseWriter, body []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
