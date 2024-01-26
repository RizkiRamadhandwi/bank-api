package usecase

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/mock/service_mock"
	"bank-api/mock/usecase_mock"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUsecaseTestSuite struct {
	suite.Suite
	jsm *service_mock.JwtServiceMock
	uum *usecase_mock.UserUsecaseMock
	auc AuthUseCase
}

func (suite *AuthUsecaseTestSuite) SetupTest() {
	suite.uum = new(usecase_mock.UserUsecaseMock)
	suite.jsm = new(service_mock.JwtServiceMock)
	suite.auc = NewAuthUseCase(suite.uum, suite.jsm)
}

var mockLogin = dto.AuthRequestDto{
	Username: "Dwi",
	Password: "Ramadhan",
}

var mockToken = dto.AuthResponseDto{
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiYW5rIiwiZXhwIjoxNzA2MTkxNTA5LCJpYXQiOjE3MDYxODU1MDksInVzZXJJZCI6IjEifQ.rMZ62VsKXpEk5PH3stzpsqZunu-q0cLSuqGCLIooqaE",
}

func (suite *AuthUsecaseTestSuite) TestLogin_Success() {
	suite.uum.On("FindUserForLogin", mockLogin.Username, mockLogin.Password).Return(mockUser, nil)
	suite.jsm.On("CreateToken", mockUser).Return(mockToken, nil)

	_, err := suite.auc.Login(mockLogin)
	assert.NoError(suite.T(), err)
}

func (suite *AuthUsecaseTestSuite) TestLogin_Fail() {
	suite.uum.On("FindUserForLogin", mockLogin.Username, mockLogin.Password).Return(entity.User{}, fmt.Errorf("error"))

	_, err := suite.auc.Login(mockLogin)
	assert.Error(suite.T(), err)
}

func (suite *AuthUsecaseTestSuite) TestLogin_CreateTokenFail() {
	suite.uum.On("FindUserForLogin", mockLogin.Username, mockLogin.Password).Return(mockUser, nil)
	suite.jsm.On("CreateToken", mockUser).Return(mockToken, fmt.Errorf("error"))

	_, err := suite.auc.Login(mockLogin)
	assert.Error(suite.T(), err)
}

func (suite *AuthUsecaseTestSuite) TestLogout_Success() {
	suite.jsm.On("DeleteToken", mockToken).Return(nil)

	err := suite.auc.Logout(mockToken)
	assert.NoError(suite.T(), err)
}

func (suite *AuthUsecaseTestSuite) TestLogout_Fail() {
	suite.jsm.On("DeleteToken", mockToken).Return(fmt.Errorf("error"))

	err := suite.auc.Logout(mockToken)
	assert.Error(suite.T(), err)
}

func TestAuthUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}
