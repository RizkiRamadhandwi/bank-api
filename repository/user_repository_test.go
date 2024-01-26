package repository

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	repo UserRepository
}

var filePathCust = "data/customers.json"
var mockFilePathCust = "../mock/data_mock/customers.json"

func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.repo = NewUserRepository(filePathCust)
}

func copyFile(sourcePath, destinationPath string) {
	sourceFile, err := os.ReadFile(sourcePath)
	if err != nil {
		log.Fatalf("Error reading source file: %v", err)
	}

	err = os.WriteFile(destinationPath, sourceFile, 0644)
	if err != nil {
		log.Fatalf("Error writing to destination file: %v", err)
	}
}

func (suite *UserRepositoryTestSuite) TestGetByIdCust_Success() {
	suite.repo = NewUserRepository(filePathCust)

	_, err := os.ReadFile(filePathCust)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdCust("1")
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetByIdCust_InvalidSuccess() {
	_, err := os.ReadFile(filePathCust)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdCust("3")
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetByIdCust_ReadDataFail() {
	filePath := "path/to/your/nonexistentfile.json"
	suite.repo = NewUserRepository(filePath)

	_, err := os.ReadFile(filePathCust)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdCust("1")
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetByIdCust_UnmarshalFail() {
	copyFile(filePathCust, mockFilePathCust)

	fileContent := `{"id":"1","Username": "user1", "Password": "pass1", "Role": "Customer", "InvalidField": "invalid"}`

	err := os.WriteFile(filePathCust, []byte(fileContent), 0644)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdCust("1")
	assert.Error(suite.T(), err)

	copyFile(mockFilePathCust, filePathCust)
}

func (suite *UserRepositoryTestSuite) TestGetForLogin_Success() {
	_, err := os.ReadFile(filePathCust)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetForLogin("customer1", "password123")
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetForLogin_InvalidSuccess() {
	_, err := os.ReadFile(filePathCust)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetForLogin("customer12345", "password12345")
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetForLogin_ReadDataFail() {
	filePath := "path/to/your/nonexistentfile.json"
	suite.repo = NewUserRepository(filePath)

	_, err := os.ReadFile(filePathCust)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetForLogin("user1", "pass1")
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestGetForLogin_UnmarshalFail() {
	copyFile(filePathCust, mockFilePathCust)

	fileContent := `{"id":"1","Username": "user1", "Password": "pass1", "Role": "Customer", "InvalidField": "invalid"}`

	err := os.WriteFile(filePathCust, []byte(fileContent), 0644)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetForLogin("user1", "pass1")
	assert.Error(suite.T(), err)

	copyFile(mockFilePathCust, filePathCust)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
