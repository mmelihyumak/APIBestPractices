package repository

import (
    "database/sql"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *User) error {
    query := "INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, ?)"
    _, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.CreatedAt)
    return err
}

func (r *UserRepository) GetUserByID(id int64) (*User, error) {
    var user User
    query := "SELECT id, name, email, password, created_at FROM users WHERE id = ?"
    err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return &user, err
}

func (r *UserRepository) UpdateUser(user *User) error {
    query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
    _, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.ID)
    return err
}

func (r *UserRepository) DeleteUser(id int64) error {
    query := "DELETE FROM users WHERE id = ?"
    _, err := r.db.Exec(query, id)
    return err
}