package usecase

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/repository"
	"bank-api/shared/model"
)

type TransactionUseCase interface {
	RegisterNewTransaction(payload entity.Transaction) (dto.TransactionDto, error)
	FindAllTransactionByCustomer(page, size int, user string) ([]dto.TransactionDto, model.Paging, error)
}

type transactionUseCase struct {
	repo repository.TransactionRepository
}

func (tu *transactionUseCase) RegisterNewTransaction(payload entity.Transaction) (dto.TransactionDto, error) {
	return tu.repo.Create(payload)
}

func (tu *transactionUseCase) FindAllTransactionByCustomer(page, size int, user string) ([]dto.TransactionDto, model.Paging, error) {
	return tu.repo.List(page, size, user)
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repo: repo}
}
