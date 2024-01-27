package controller

import (
	"bank-api/entity/dto"
	"bank-api/mock/service_mock"
	"bank-api/mock/usecase_mock"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *usecase_mock.AuthUsecaseMock
	js  *service_mock.JwtServiceMock
	ac  *AuthController
}

func (suite *AuthControllerTestSuite) SetupTest() {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := router.Group("/api/v1")
	suite.rg = rg
	suite.aum = new(usecase_mock.AuthUsecaseMock)
	suite.js = new(service_mock.JwtServiceMock)
	suite.ac = NewAuthController(suite.aum, suite.rg, suite.js)
	suite.ac.Route()
}

var mockLogin = dto.AuthRequestDto{
	Username: "Dwi",
	Password: "Ramadhan",
}

var mockToken = dto.AuthResponseDto{
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiYW5rIiwiZXhwIjoxNzA2MTkxNTA5LCJpYXQiOjE3MDYxODU1MDksInVzZXJJZCI6IjEifQ.rMZ62VsKXpEk5PH3stzpsqZunu-q0cLSuqGCLIooqaE",
}

func (suite *AuthControllerTestSuite) TestLoginHandler_Success() {
	suite.aum.On("Login", mockLogin).Return(mockToken, nil)

	requestBody, _ := json.Marshal(mockLogin)
	request, err := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	suite.ac.loginHandler(ctx)

	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *AuthControllerTestSuite) TestLoginHandler_BadRequest() {
	mockLogin := dto.AuthRequestDto{}
	mockError := errors.New("example error message")

	suite.aum.On("Login", &mockLogin).Return(mockToken, mockError)

	request, err := http.NewRequest(http.MethodPost, "/auth/login", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	suite.ac.loginHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *AuthControllerTestSuite) TestLoginHandler_InternalServerError() {
	mockError := errors.New("example error message")
	suite.aum.On("Login", mockLogin).Return(mockToken, mockError)

	requestBody, _ := json.Marshal(mockLogin)
	request, err := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	suite.ac.loginHandler(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

// error
// func (suite *AuthControllerTestSuite) TestLogoutHandler_Success() {
// 	suite.js.On("ParseToken", mockToken.Token).Return(jwt.MapClaims{"key": "value"}, nil)
// 	suite.aum.On("Logout", dto.AuthResponseDto{Token: mockToken.Token}).Return(nil)

// 	requestHeader, _ := json.Marshal(mockToken)
// 	request, err := http.NewRequest(http.MethodPost, "/auth/logout", bytes.NewReader(requestHeader))
// 	assert.NoError(suite.T(), err)
// 	request.Header.Set("Authorization", "Bearer "+mockToken.Token)

// 	responseRecorder := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(responseRecorder)
// 	ctx.Request = request

// 	suite.ac.logoutHandler(ctx)

// 	assert.Equal(suite.T(), http.StatusNoContent, responseRecorder.Code)
// }

func (suite *AuthControllerTestSuite) TestLogoutHandler_GetHeaderFail() {
	suite.aum.On("Logout", mockToken.Token).Return(jwt.MapClaims{"key": "value"}, nil)

	requestHeader, _ := json.Marshal(mockToken)
	request, err := http.NewRequest(http.MethodPost, "/auth/logout", bytes.NewReader(requestHeader))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	suite.ac.logoutHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *AuthControllerTestSuite) TestLogoutHandler_InvalidTokenFail() {
	suite.js.On("ParseToken", "invalidTokenFormat").Return(jwt.MapClaims{"key": "value"}, errors.New("Invalid token format"))

	request, err := http.NewRequest(http.MethodPost, "/auth/logout", nil)
	assert.NoError(suite.T(), err)
	request.Header.Set("Authorization", "Bearer invalidTokenFormat")

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	suite.ac.logoutHandler(ctx)

	assert.Equal(suite.T(), http.StatusUnauthorized, responseRecorder.Code)
}

func (suite *AuthControllerTestSuite) TestLogoutHandler_LogoutError() {
	// Your test setup
	mockError := errors.New("logout failed")
	suite.js.On("ParseToken", "mock").Return(jwt.MapClaims{"key": "value"}, nil)
	suite.aum.On("Logout", dto.AuthResponseDto{Token: "mock"}).Return(mockError)

	request, err := http.NewRequest(http.MethodPost, "/auth/logout", nil)
	assert.NoError(suite.T(), err)
	request.Header.Set("Authorization", "Bearer mock")

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	suite.ac.logoutHandler(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
	assert.Contains(suite.T(), responseRecorder.Body.String(), mockError.Error())
}

func (suite *AuthControllerTestSuite) TestLogoutHandler_TokenAfterLogoutFail() {
	// Your test setup
	suite.js.On("ParseToken", "mock").Return(jwt.MapClaims{"key": "value"}, nil)
	suite.aum.On("Logout", dto.AuthResponseDto{Token: "mock"}).Return(nil)

	request, err := http.NewRequest(http.MethodPost, "/auth/logout", nil)
	assert.NoError(suite.T(), err)
	request.Header.Set("Authorization", "Bearer mock")

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	suite.ac.logoutHandler(ctx)

	assert.Equal(suite.T(), http.StatusUnauthorized, responseRecorder.Code)
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
