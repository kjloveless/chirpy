package main

import (
    "net/http"
)

func main() {
    sm := http.NewServeMux()
    sm.Handle("/", http.FileServer(http.Dir(".")))
    server := &http.Server{
        Addr:       ":8080",
        Handler:    sm,
    }

    server.ListenAndServe()
}
