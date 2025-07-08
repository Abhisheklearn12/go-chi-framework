
package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World, how are you!"))
    })
    http.ListenAndServe(":3000", r)
}
// This is just for Basic Golang project.

/*
* This is just for fun and learning backends.
* Try to learn.
*/
