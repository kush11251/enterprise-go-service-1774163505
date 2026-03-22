package service

import (
    "enterprise-go-app/pkg/models"
    "enterprise-go-app/pkg/repository"
)

// UserService business behavior
type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) ListUsers() []models.User {
    return s.repo.List()
}

func (s *UserService) GetUser(id int) (*models.User, bool) {
    return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(username, email string) models.User {
    return s.repo.Create(username, email)
}