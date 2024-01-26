package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MerchantRepositoryTestSuite struct {
	suite.Suite
	repo MerchantRepository
}

var filePathMerc = "data/merchants.json"
var mockFilePathMerc = "../mock/data_mock/merchants.json"

func (suite *MerchantRepositoryTestSuite) SetupTest() {
	suite.repo = NewMerchantRepository(filePathMerc)
}

func (suite *MerchantRepositoryTestSuite) TestGetByIdMerc_Success() {
	_, err := os.ReadFile(filePathMerc)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdMerc("1")
	assert.NoError(suite.T(), err)
}

func (suite *MerchantRepositoryTestSuite) TestGetByIdMerc_InvalidSuccess() {
	_, err := os.ReadFile(filePathCust)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdMerc("100")
	assert.NoError(suite.T(), err)
}

func (suite *MerchantRepositoryTestSuite) TestGetByIdMerc_ReadDataFail() {
	filePath := "path/to/your/nonexistentfile.json"
	suite.repo = NewMerchantRepository(filePath)

	_, err := os.ReadFile(filePathMerc)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdMerc("1")
	assert.Error(suite.T(), err)
}

func (suite *MerchantRepositoryTestSuite) TestGetByIdMerc_UnmarshalFail() {
	copyFile(filePathMerc, mockFilePathMerc)

	fileContent := `{"id":"1","Username": "user1", "Password": "pass1", "Role": "Customer", "InvalidField": "invalid"}`

	err := os.WriteFile(filePathMerc, []byte(fileContent), 0644)
	assert.NoError(suite.T(), err)

	_, err = suite.repo.GetByIdMerc("1")
	assert.Error(suite.T(), err)

	copyFile(mockFilePathMerc, filePathMerc)
}

func TestMerchantRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MerchantRepositoryTestSuite))
}
