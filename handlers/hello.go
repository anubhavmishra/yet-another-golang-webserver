package handlers

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

type helloWorldHandler struct {
	Message string
	Version string
}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: h.Message,
		Version: h.Version,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func HelloWorldHandler(message string, version string) http.Handler {
	if message == "" {
		message = "Hello World!"
	}

	return &helloWorldHandler{
		Message: message,
		Version: version,
	}
}
