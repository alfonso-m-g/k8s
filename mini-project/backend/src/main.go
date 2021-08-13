package main

import (
	"encoding/json"
    "net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time	   time.Time  `json:"time"`
	HostName   string     `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	resp := HandsOn{
		Time: time.Now(),
		HostName: os.Getenv("HOSTNAME"),
	}

	jsonResp, err := json.Marshal(&resp)
	
	if err != nil {
		w.Write([]byte("Error"))
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(jsonResp)
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":8080", nil)
}
