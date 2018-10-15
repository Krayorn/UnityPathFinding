package main

import (
    "net/http"
    "io"
    "os"
)

func main() {
    port := os.Getenv("PORT")

    http.HandleFunc("/", sayHello)

    http.ListenAndServe(":"+port, nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello")
}
