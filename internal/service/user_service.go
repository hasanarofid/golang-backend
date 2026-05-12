package service

import (
	"errors"

	"github.com/hasanarofid/golang-backend/internal/model"
	"github.com/hasanarofid/golang-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Register(user model.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hash)

	return s.Repo.Create(user)
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	user, _ := s.Repo.FindByEmail(email)

	if user == nil {
		return nil, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
