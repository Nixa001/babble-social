package service

import (
	"backend/models"
	r "backend/server/repositories"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo r.UserRepository
	SessRepo r.SessionRepository
}

func (a *AuthService) init() {
	a.UserRepo = *r.UserRepo
	a.SessRepo = *r.SessionRepo
}

func (a *AuthService) CreateUser(user *models.User) error {
	_, err := a.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	err = a.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) CheckCredentials(email, password string) (*models.User, error) {
	user, err := a.UserRepo.GetUserByEmail(email)
	if err != nil {
		return &models.User{}, fmt.Errorf("Invalid Credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &models.User{}, fmt.Errorf("Invalid Credentials")
	}
	return user, nil
}
