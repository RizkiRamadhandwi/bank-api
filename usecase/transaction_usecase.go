package usecase

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/repository"
	"bank-api/shared/model"
	"fmt"
)

type TransactionUseCase interface {
	RegisterNewTransaction(payload entity.Transaction) (dto.TransactionDto, error)
	FindAllTransactionByCustomer(page, size int, user string) ([]dto.TransactionDto, model.Paging, error)
}

type transactionUseCase struct {
	repo repository.TransactionRepository
}

func (tu *transactionUseCase) RegisterNewTransaction(payload entity.Transaction) (dto.TransactionDto, error) {
	if payload.MerchantID == "" || payload.Amount <= 0 {
		return dto.TransactionDto{}, fmt.Errorf("oops, field required")
	}

	transaction, err := tu.repo.Create(payload)
	if err != nil {
		return dto.TransactionDto{}, fmt.Errorf("merchant id not found")
	}

	return transaction, nil
}

func (tu *transactionUseCase) FindAllTransactionByCustomer(page, size int, user string) ([]dto.TransactionDto, model.Paging, error) {
	return tu.repo.List(page, size, user)
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repo: repo}
}
