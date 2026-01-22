package main

import (
    "fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "socialnet skeleton: Assignment 3")
    })

    http.ListenAndServe(":8080", mux)
}
