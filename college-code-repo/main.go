package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/sangan007/CodeRepo/backend/internal/auth"
    "github.com/sangan007/CodeRepo/backend/internal/user"
    "github.com/sangan007/CodeRepo/backend/internal/repository"
)



func main() {
    r := mux.NewRouter()

    // Authentication endpoints
    r.HandleFunc("/api/login", auth.Login).Methods("POST")

    // User management endpoints
    r.HandleFunc("/api/signup", user.SignUp).Methods("POST")

    // Repository endpoints
    r.HandleFunc("/api/repositories", repository.CreateRepository).Methods("POST")

    // Static file serving
    fs := http.FileServer(http.Dir("./frontend"))
    r.PathPrefix("/").Handler(fs)

    // Start the server
    http.ListenAndServe(":8000", r)
}
