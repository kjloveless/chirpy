package main

import (
    "fmt"
    "net/http"
)

type apiConfig struct {
    fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cfg.fileserverHits += 1
        next.ServeHTTP(w, r)
    }) 
}

func (cfg *apiConfig) getFileserverHits() int {
    return cfg.fileserverHits
}

func (cfg *apiConfig) resetFileserverHits() {
    cfg.fileserverHits = 0
}

func main() {
    sm := http.NewServeMux()
    cfg := apiConfig{}
    handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
    sm.Handle("/app/*", cfg.middlewareMetricsInc(handler))

    sm.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(200)
        w.Write([]byte("OK"))
    })

    sm.HandleFunc("GET /metrics", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(200)
        w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.getFileserverHits())))
    })

    sm.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(200)
        cfg.resetFileserverHits()
    })
    server := &http.Server{
        Addr:       ":8080",
        Handler:    sm,
    }

    server.ListenAndServe()
}
