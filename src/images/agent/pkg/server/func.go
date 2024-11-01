package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	log.Printf("FC Initialize End RequestId: %s", requestID)
	fmt.Fprintf(w, "Function is initialized, request_id: %s\n", requestID)
}

func Invoke(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("x-fc-request-id")
	log.Printf("FC Invoke Start RequestId: %s", requestID)

	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var eventData EventData
	// 将请求体解析为 JSON
	if err := json.Unmarshal(body, &eventData); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	//// 获取任务类型和图像 URL
	//taskType := eventData.Type
	//imageURL := eventData.ImageURL
	//
	//// 返回响应
	response := map[string]string{"image_url": "newImageURL"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create response JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
