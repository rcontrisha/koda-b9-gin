package service

import (
	"errors"
	"rcontrisha/koda-b9-gin/internal/dto"
)

var Users = []dto.User{}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) LoginService(payload dto.User) error {
	if len(Users) > 0 {
		for _, user := range Users {
			if payload.Username == user.Username && payload.Password == user.Password {
				return nil
			}
		}
	}

	return errors.New("Login failed. Invalid username or password.")
}

func (a *AuthService) RegisterService(payload dto.User) error {
	if len(payload.Password) < 8 {
		return errors.New("Register failed. Password must at least 8 characters.")
	}

	if len(Users) > 0 {
		for _, user := range Users {
			if payload.Username == user.Username {
				return errors.New("Register failed. User already exist.")
			}
		}
	}

	Users = append(Users, payload)
	
	return nil
}
