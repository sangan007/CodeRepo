package auth

import (
    "encoding/json"
    "net/http"

    "github.com/sangan007/CodeRepo/backend/internal/model"
    "github.com/sangan007/CodeRepo/backend/internal/repository"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var user model.User
    json.NewDecoder(r.Body).Decode(&user)

    valid, err := repository.ValidateUser(user)
    if err != nil || !valid {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
}
