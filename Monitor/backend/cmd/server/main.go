package main

import (
    "log"
    "net/http"

    "jms.smartmonitor/api/handlers"
)

func main() {
    r := handlers.InitRouter()

    log.Println("Smart Network Monitor backend starting on :8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("server error:", err)
    }
}