package service

import (
	"backend/models"
	r "backend/server/repositories"
	"backend/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

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

func (a *AuthService) CreateUser(user models.FormatedUser) error {
	_, err := a.UserRepo.GetUserByEmail(user.Email)
	if err == nil {
		return fmt.Errorf("this email is already in use")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	err = a.UserRepo.SaveUser(user)
	return err
}

func (a *AuthService) CheckCredentials(email, password string) (models.User, error) {
	user, err := a.UserRepo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid Email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, fmt.Errorf("invalid Password")
	}
	return user, nil
}

func (a *AuthService) VerifyToken(r *http.Request) (session models.Session, err error) {
	token := r.Header.Get("Authorization")
	fmt.Println("VerifyToken u8u9i9u8", token)

	if token == "" {
		token, err = url.QueryUnescape(r.URL.Query().Get("token"))
		if err != nil {
			fmt.Println("Erreur token", err.Error())
			return models.Session{}, fmt.Errorf("invalid token")
		}
		token = strings.ReplaceAll(token, " ", "+")
	}

	if token == "" {
		return models.Session{}, fmt.Errorf("missing token")
	}
	session, err = a.SessRepo.GetSession(token)
	if err != nil {
		fmt.Println("Error getting session", err)
		return models.Session{}, err
	}
	if session.Expiration.Before(time.Now()) {
		fmt.Println("token expired")
		return models.Session{}, fmt.Errorf("token expired")
	}
	return session, nil
}

func (a *AuthService) RemoveSession(token string) error {
	return a.SessRepo.DeleteSession(token)
}

func (a *AuthService) RemExistingSession(userId int) error {
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

func (a *AuthService) CreateSession(user models.User) (models.Session, error) {
	err := a.RemExistingSession(user.Id)
	if err != nil {
		log.Println("Error removing existing session", err)
		return models.Session{}, err
	}
	token := utils.GenerateToken()
	expiration := time.Now().Add(3 * time.Hour)
	err = a.SessRepo.SaveSession(models.Session{User_id: user.Id, Token: token, Expiration: expiration})
	if err != nil {
		return models.Session{}, err
	}
	return models.Session{User_id: user.Id, Token: token, Expiration: expiration}, nil

}
