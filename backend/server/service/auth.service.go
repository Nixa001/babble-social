package service

import (
	"backend/models"
	r "backend/server/repositories"
	"fmt"
	"net/http"
	"strings"

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
		return &models.User{}, fmt.Errorf("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &models.User{}, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

func (a *AuthService) VerifyToken(r *http.Request) (session models.Session, err error) {
	token := r.Header.Get("Authorization")
	if strings.TrimSpace(token) == "" {
		return models.Session{}, fmt.Errorf("no token provided")
	}
	return session, nil
}

func (a *AuthService) RemoveSession(token string) error {
	return a.SessRepo.DeleteSession(token)
}

func (a *AuthService) RemExistingSession(userId string) error {
	sessions, err := a.SessRepo.GetSessionByUserId(userId)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		err = a.SessRepo.DeleteSession(session.Token)
		if err != nil {
			return err
		}
	}
	return nil
}
