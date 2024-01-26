package repository

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/shared/model"
	"encoding/json"
	"log"
	"math"
	"os"
	"time"
)

type TransactionRepository interface {
	Create(payload entity.Transaction) (dto.TransactionDto, error)
	List(page, size int, user string) ([]dto.TransactionDto, model.Paging, error)
}

type transactionRepository struct {
	filePath  string
	userRepo  UserRepository
	merchRepo MerchantRepository
}

func (tr *transactionRepository) Create(payload entity.Transaction) (dto.TransactionDto, error) {
	readData, err := os.ReadFile(tr.filePath)
	if err != nil {
		log.Printf("TransactionRepository.Create: %v \n", err.Error())
		return dto.TransactionDto{}, err
	}

	var transactions []dto.TransactionDto
	err = json.Unmarshal(readData, &transactions)
	if err != nil {
		log.Printf("TransactionRepository.Create: %v \n", err.Error())
		return dto.TransactionDto{}, err
	}

	customer, err := tr.userRepo.GetByIdCust(payload.UserID)
	if err != nil {
		log.Printf("TransactionRepository.Create: %v \n", err.Error())
		return dto.TransactionDto{}, err
	}

	merchant, err := tr.merchRepo.GetByIdMerc(payload.MerchantID)
	if err != nil {
		log.Printf("TransactionRepository.Create: %v \n", err.Error())
		return dto.TransactionDto{}, err
	}

	newTransaction := dto.TransactionDto{
		ID:        len(transactions) + 1,
		Customer:  customer,
		Merchant:  merchant,
		Amount:    payload.Amount,
		CreatedAt: time.Now(),
	}

	transactions = append(transactions, newTransaction)

	newData, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		log.Printf("TransactionRepository.Create: %v \n", err.Error())
		return dto.TransactionDto{}, err
	}

	err = os.WriteFile(tr.filePath, newData, 0644)
	if err != nil {
		log.Printf("TransactionRepository.Create: %v \n", err.Error())
		return dto.TransactionDto{}, err
	}

	return newTransaction, nil
}

func (tr *transactionRepository) List(page, size int, user string) ([]dto.TransactionDto, model.Paging, error) {
	var transactions []dto.TransactionDto

	readData, err := os.ReadFile(tr.filePath)
	if err != nil {
		log.Printf("TransactionRepository.List: %v \n", err.Error())
		return nil, model.Paging{}, err
	}

	err = json.Unmarshal(readData, &transactions)
	if err != nil {
		log.Printf("TransactionRepository.List: %v \n", err.Error())
		return nil, model.Paging{}, err
	}

	userTransactions := make([]dto.TransactionDto, 0)
	for _, transaction := range transactions {
		if transaction.Customer.ID == user {
			userTransactions = append(userTransactions, transaction)
		}
	}

	totalRows := len(userTransactions)
	totalPages := int(math.Ceil(float64(totalRows) / float64(size)))

	startIndex := (page - 1) * size
	endIndex := startIndex + size

	if endIndex > totalRows {
		endIndex = totalRows
	}

	pageTransactions := userTransactions[startIndex:endIndex]

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  totalPages,
	}

	return pageTransactions, paging, nil
}

func NewTransactionRepository(filePath string, userRepo UserRepository, merchRepo MerchantRepository) TransactionRepository {
	return &transactionRepository{filePath: filePath, userRepo: userRepo, merchRepo: merchRepo}
}
