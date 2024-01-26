package repository_mock

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/shared/model"

	"github.com/stretchr/testify/mock"
)

type TransactionRepoMock struct {
	mock.Mock
}

func (t *TransactionRepoMock) Create(payload entity.Transaction) (dto.TransactionDto, error) {
	args := t.Called(payload)
	return args.Get(0).(dto.TransactionDto), args.Error(1)
}

func (t *TransactionRepoMock) List(page, size int, user string) ([]dto.TransactionDto, model.Paging, error) {
	args := t.Called(page, size, user)
	return args.Get(0).([]dto.TransactionDto), args.Get(1).(model.Paging), args.Error(2)
}
