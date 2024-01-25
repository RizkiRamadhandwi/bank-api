package usecase

import (
	"bank-api/entity"
	"bank-api/repository"
	"errors"
	"fmt"
)

type UserUseCase interface {
	FindUserForLogin(username, password string) (entity.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (uu *userUseCase) FindUserForLogin(username, password string) (entity.User, error) {

	if username == "" && password == "" {
		return entity.User{}, errors.New("username and password required")
	}
	if username == "" {
		return entity.User{}, errors.New("username required")
	}
	if password == "" {
		return entity.User{}, errors.New("password required")
	}

	userExist, err := uu.repo.GetForLogin(username, password)
	if err != nil {
		return entity.User{}, fmt.Errorf("user doesn't exists")
	}

	return userExist, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
