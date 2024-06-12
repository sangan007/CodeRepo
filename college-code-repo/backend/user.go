package user

import (
    "encoding/json"
    "net/http"

    "github.com/sangan007/CodeRepo/backend/internal/model"
    "github.com/sangan007/CodeRepo/backend/internal/repository"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
    var user model.User
    json.NewDecoder(r.Body).Decode(&user)

    err := repository.CreateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
