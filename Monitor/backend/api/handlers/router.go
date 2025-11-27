package handlers

import (
    "net/http"

    "github.com/gorilla/mux"
)

// InitRouter initializes all API routes.
// This file is intentionally minimal – real handlers will be added gradually
// according to the Block specification (Auth, Settings, Discovery, Unified, etc.).
func InitRouter() *mux.Router {
    r := mux.NewRouter()

    // Health check
    r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte(`{"status":"ok"}`))
    }).Methods(http.MethodGet)

    // TODO: add subrouters for:
    // - /api/auth
    // - /api/settings
    // - /api/discovery
    // - /api/unified
    // - /api/ping
    // - /api/diagnostics
    // - /api/map
    // - /api/events
    // - /api/backup
    // - /api/system

    return r
}