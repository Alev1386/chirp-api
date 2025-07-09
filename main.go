package main


import (
    "encoding/json"
    "net/http"
)

type ChirpParam struct {
    ID       int     `json:"id"`
    StartHz  float64 `json:"start_hz"`
    EndHz    float64 `json:"end_hz"`
    Duration float64 `json:"duration_us"`
}

var chirps = []ChirpParam{
    {ID: 1, StartHz: 100e6, EndHz: 200e6, Duration: 10},
    {ID: 2, StartHz: 150e6, EndHz: 250e6, Duration: 15},
}

func chirpsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(chirps)
}


func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/chirps", chirpsHandler)
    http.ListenAndServe(":8080", nil)
}
