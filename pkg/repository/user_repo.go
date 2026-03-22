package repository

import "enterprise-go-app/pkg/models"

// UserRepository is an in-memory storage implementation
type UserRepository struct {
    users []models.User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{users: []models.User{{ID: 1, Username: "alpha", Email: "alpha@example.com", Active: true}}}
}

func (r *UserRepository) List() []models.User {
    return append([]models.User(nil), r.users...)
}

func (r *UserRepository) GetByID(id int) (*models.User, bool) {
    for _, u := range r.users {
        if u.ID == id {
            return &u, true
        }
    }
    return nil, false
}

func (r *UserRepository) Create(username, email string) models.User {
    newID := len(r.users) + 1
    user := models.User{ID: newID, Username: username, Email: email, Active: true}
    r.users = append(r.users, user)
    return user
}