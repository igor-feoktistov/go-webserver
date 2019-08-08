package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "sort"
)

func handler(w http.ResponseWriter, r *http.Request) {
    var keys []string

    host, _ := os.Hostname()
    for k := range r.Header {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    log.Printf("Serving request: %s", r.URL.Path)
    fmt.Fprintf(w, "<p>GoLang WebServer, version: 1.0.0</p>\n")
    fmt.Fprintf(w, "<p><b>Hostname:</b> %s</p>\n", host)
    fmt.Fprintf(w, "<p><b>Request Headers: </b></p>\n")
    for _, k := range keys {
        fmt.Fprintf(w, "<b>%s:</b> %s</br>\n", k, r.Header[k])
    }
}

func main() {
    // PORT environment variable, default is 8080
    port := "8080"
    if fromEnv := os.Getenv("PORT"); fromEnv != "" {
	port = fromEnv
    }
    server := http.NewServeMux()
    server.HandleFunc("/", handler)
    log.Printf("Server listening on port %s", port)
    err := http.ListenAndServe(":" + port, server)
    log.Fatal(err)
}
