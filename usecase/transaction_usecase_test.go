package usecase

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/mock/repository_mock"
	"bank-api/shared/model"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionUsecaseTestSuite struct {
	suite.Suite
	trm *repository_mock.TransactionRepoMock
	tuc TransactionUseCase
}

func (suite *TransactionUsecaseTestSuite) SetupTest() {
	suite.trm = new(repository_mock.TransactionRepoMock)
	suite.tuc = NewTransactionUseCase(suite.trm)
}

// RegisterNewTransaction(payload entity.Transaction) (dto.TransactionDto, error)

var mockTransaction = entity.Transaction{
	ID:         1,
	UserID:     "1",
	MerchantID: "1",
	Amount:     50000,
	CreatedAt:  time.Now(),
}

var mockTransactionDto = dto.TransactionDto{
	ID: 1,
	Customer: dto.UserDto{
		ID:   "1",
		Name: "John Doe",
	},
	Merchant: dto.MerchantDto{
		ID:   "1",
		Name: "Google",
	},
	Amount:    50000,
	CreatedAt: time.Now(),
}

func (suite *TransactionUsecaseTestSuite) TestFindAllTransactionByCustomer_Success() {
	suite.trm.On("List", 1, 5, "1").Return([]dto.TransactionDto{}, model.Paging{}, nil)

	_, _, err := suite.tuc.FindAllTransactionByCustomer(1, 5, "1")
	assert.NoError(suite.T(), err)
}

func (suite *TransactionUsecaseTestSuite) TestRegisterNewTransaction_Success() {
	suite.trm.On("Create", mockTransaction).Return(mockTransactionDto, nil)

	_, err := suite.tuc.RegisterNewTransaction(mockTransaction)
	assert.NoError(suite.T(), err)
}

func (suite *TransactionUsecaseTestSuite) TestRegisterNewTransaction_NoFieldFail() {
	mockTransaction.MerchantID = ""
	suite.trm.On("Create", mockTransaction).Return(dto.TransactionDto{}, fmt.Errorf("oops, field required"))

	_, err := suite.tuc.RegisterNewTransaction(mockTransaction)
	assert.Error(suite.T(), err)
}

func (suite *TransactionUsecaseTestSuite) TestRegisterNewTransaction_NotFoundMerchantIdFail() {
	mockTransaction.MerchantID = "5"
	suite.trm.On("Create", mockTransaction).Return(dto.TransactionDto{}, fmt.Errorf("merchant id not found"))

	_, err := suite.tuc.RegisterNewTransaction(mockTransaction)
	assert.Error(suite.T(), err)
}

func TestTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionUsecaseTestSuite))
}
