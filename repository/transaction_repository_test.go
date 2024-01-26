package repository

import (
	"bank-api/entity"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionRepositoryTestSuite struct {
	suite.Suite
	repo      TransactionRepository
	userRepo  UserRepository
	merchRepo MerchantRepository
}

var filePathTrans = "data/transactions.json"
var mockFilePathTrans = "../mock/data_mock/transactions.json"
var mockCreate = entity.Transaction{
	UserID:     "1",
	MerchantID: "1",
	Amount:     1000000,
}

func (suite *TransactionRepositoryTestSuite) SetupTest() {
	suite.userRepo = NewUserRepository(filePathCust)
	suite.merchRepo = NewMerchantRepository(filePathMerc)
	suite.repo = NewTransactionRepository(filePathTrans, suite.userRepo, suite.merchRepo)
}

func (suite *TransactionRepositoryTestSuite) TestCreate_Success() {
	copyFile(filePathTrans, mockFilePathTrans)
	_, err := os.ReadFile(filePathTrans)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.Create(mockCreate)
	assert.NoError(suite.T(), err)

	copyFile(mockFilePathTrans, filePathTrans)
}

func (suite *TransactionRepositoryTestSuite) TestCreate_ReadDataFail() {
	filePath := "path/to/your/nonexistentfile.json"
	suite.userRepo = NewUserRepository(filePathCust)
	suite.merchRepo = NewMerchantRepository(filePathMerc)
	suite.repo = NewTransactionRepository(filePath, suite.userRepo, suite.merchRepo)

	_, err := os.ReadFile(filePathTrans)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.Create(mockCreate)
	assert.Error(suite.T(), err)
}

func (suite *TransactionRepositoryTestSuite) TestCreate_UnmarshalFail() {
	copyFile(filePathTrans, mockFilePathTrans)

	fileContent := `{"id":"1","Username": "user1", "Password": "pass1", "Role": "Customer", "InvalidField": "invalid"}`

	err := os.WriteFile(filePathTrans, []byte(fileContent), 0644)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.Create(mockCreate)
	assert.Error(suite.T(), err)

	copyFile(mockFilePathTrans, filePathTrans)
}

func (suite *TransactionRepositoryTestSuite) TestCreate_GetUserFail() {
	suite.userRepo = NewUserRepository("1.json")
	suite.merchRepo = NewMerchantRepository(filePathMerc)
	suite.repo = NewTransactionRepository(filePathTrans, suite.userRepo, suite.merchRepo)

	_, err := os.ReadFile(filePathTrans)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.Create(mockCreate)
	assert.Error(suite.T(), err)
}

func (suite *TransactionRepositoryTestSuite) TestCreate_GetMerchantFail() {
	suite.userRepo = NewUserRepository(filePathCust)
	suite.merchRepo = NewMerchantRepository("1.json")
	suite.repo = NewTransactionRepository(filePathTrans, suite.userRepo, suite.merchRepo)

	_, err := os.ReadFile(filePathTrans)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.Create(mockCreate)
	assert.Error(suite.T(), err)
}

func (suite *TransactionRepositoryTestSuite) TestList_Success() {
	_, err := os.ReadFile(filePathTrans)
	assert.NoError(suite.T(), err)

	_, _, err = suite.repo.List(1, 10, "1")
	assert.NoError(suite.T(), err)
}

func (suite *TransactionRepositoryTestSuite) TestList_ReadDataFail() {
	filePath := "path/to/your/nonexistentfile.json"
	suite.userRepo = NewUserRepository(filePathCust)
	suite.merchRepo = NewMerchantRepository(filePathMerc)
	suite.repo = NewTransactionRepository(filePath, suite.userRepo, suite.merchRepo)

	_, err := os.ReadFile(filePathTrans)
	assert.NoError(suite.T(), err)

	_, _, err = suite.repo.List(1, 10, "1")
	assert.Error(suite.T(), err)
}

func (suite *TransactionRepositoryTestSuite) TestList_UnmarshalFail() {
	copyFile(filePathTrans, mockFilePathTrans)

	fileContent := `{"id":"1","Username": "user1", "Password": "pass1", "Role": "Customer", "InvalidField": "invalid"}`

	err := os.WriteFile(filePathTrans, []byte(fileContent), 0644)
	assert.NoError(suite.T(), err)

	_, _, err = suite.repo.List(1, 10, "1")
	assert.Error(suite.T(), err)

	copyFile(mockFilePathTrans, filePathTrans)
}

func TestTransactionsRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}
