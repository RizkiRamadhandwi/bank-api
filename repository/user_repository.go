package repository

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"encoding/json"
	"log"
	"os"
)

type UserRepository interface {
	GetForLogin(username, password string) (entity.User, error)
	GetByIdCust(id string) (dto.UserDto, error)
	GetByIdMerc(id string) (dto.MerchantDto, error)
}

type userRepository struct {
	filePath string
}

func (ur *userRepository) GetByIdCust(id string) (dto.UserDto, error) {
	readData, err := os.ReadFile(ur.filePath)
	if err != nil {
		log.Printf("UserRepository.GetByID: %v \n", err.Error())
		return dto.UserDto{}, err
	}

	var users []dto.UserDto
	err = json.Unmarshal(readData, &users)
	if err != nil {
		log.Printf("UserRepository.GetByID: %v \n", err.Error())
		return dto.UserDto{}, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return dto.UserDto{}, err
}

func (*userRepository) GetByIdMerc(id string) (dto.MerchantDto, error) {
	readData, err := os.ReadFile("repository/data/merchants.json")
	if err != nil {
		log.Printf("UserRepository.GetByID: %v \n", err.Error())
		return dto.MerchantDto{}, err
	}

	var users []dto.MerchantDto
	err = json.Unmarshal(readData, &users)
	if err != nil {
		log.Printf("UserRepository.GetByID: %v \n", err.Error())
		return dto.MerchantDto{}, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return dto.MerchantDto{}, err
}

func (ur *userRepository) GetForLogin(username, password string) (entity.User, error) {
	var users []entity.User

	readData, err := os.ReadFile(ur.filePath)
	if err != nil {
		return entity.User{}, err
	}

	err = json.Unmarshal(readData, &users)
	if err != nil {
		return entity.User{}, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return entity.User{}, err
}

func NewUserRepository(filepath string) UserRepository {
	return &userRepository{filePath: filepath}
}
