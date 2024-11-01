package server

import (
	"fmt"
	"log"
	"net/http"
)

type EventData struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
}

func Initialize(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("x-fc-request-id")
	log.Printf("FC Initialize Start RequestId: %s", requestID)
	fmt.Fprintf(w, "Function is initialized, request_id: %s\n", requestID)
}

func Invoke(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("x-fc-request-id")
	log.Printf("FC Invoke Start RequestId: %s", requestID)
	RunComfyUI(w, r)
}
