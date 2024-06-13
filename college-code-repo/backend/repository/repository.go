package repository

import (
    "database/sql"
    "github.com/sangan007/CodeRepo/backend/internal/model"
    _ "github.com/lib/pq"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "user=username password=password dbname=dbname sslmode=disable")
    if err != nil {
        panic(err)
    }
}

func CreateUser(user model.User) error {
    _, err := db.Exec("INSERT INTO users (name, student_id, password, role) VALUES ($1, $2, $3, $4)",
        user.Name, user.StudentID, user.Password, user.Role)
    return err
}

func ValidateUser(user model.User) (bool, error) {
    var dbUser model.User
    err := db.QueryRow("SELECT id, password FROM users WHERE student_id=$1 AND role=$2", user.StudentID, user.Role).Scan(&dbUser.ID, &dbUser.Password)
    if err != nil {
        return false, err
    }
    if dbUser.Password != user.Password {
        return false, nil
    }
    return true, nil
}

func CreateRepository(repo model.Repository) error {
    _, err := db.Exec("INSERT INTO repositories (name, description, owner_id) VALUES ($1, $2, $3)",
        repo.Name, repo.Description, repo.OwnerID)
    return err
}
