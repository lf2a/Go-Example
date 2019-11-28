package main

import (
    "fmt"
    "log"
    "net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL.Path)
        f(w, r)
    }
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "home")
}

func about(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "about")
}

func main() {
    http.HandleFunc("/", logging(home))
    http.HandleFunc("/about", logging(about))

    http.ListenAndServe(":8000", nil)
}
