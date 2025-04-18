package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"time": time.Now().Format(time.RFC3339),
	})
}

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Println("Service C running on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

