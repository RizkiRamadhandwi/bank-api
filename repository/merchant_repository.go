package repository

import (
	"bank-api/entity/dto"
	"encoding/json"
	"log"
	"os"
)

type MerchantRepository interface {
	GetByIdMerc(id string) (dto.MerchantDto, error)
}

type merchantRepository struct {
	filePath string
}

func (mr *merchantRepository) GetByIdMerc(id string) (dto.MerchantDto, error) {
	readData, err := os.ReadFile(mr.filePath)
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

func NewMerchantRepository(filepath string) MerchantRepository {
	return &merchantRepository{filePath: filepath}
}
