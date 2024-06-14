package main

import (
    "net/http"
)

func main() {
    sm := http.NewServeMux()
    sm.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
    sm.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(200)
        w.Write([]byte("OK"))
    })
    server := &http.Server{
        Addr:       ":8080",
        Handler:    sm,
    }

    server.ListenAndServe()
}
