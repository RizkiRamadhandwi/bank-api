package controller

import (
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/mock/middleware_mock"
	"bank-api/mock/usecase_mock"
	"bank-api/shared/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	cx  *gin.Context
	tum *usecase_mock.TransactionUsecaseMock
	amm *middleware_mock.AuthMiddlewareMock
	tc  *TransactionController
}

func (suite *TransactionControllerTestSuite) SetupTest() {

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := router.Group("/api/v1")
	suite.rg = rg
	suite.tum = new(usecase_mock.TransactionUsecaseMock)
	suite.amm = new(middleware_mock.AuthMiddlewareMock)
	suite.tc = NewTransactionController(suite.tum, suite.rg, suite.amm)
	suite.tc.Route()

	suite.cx, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.cx.Set("user", "1")
}

var mockTransaction = entity.Transaction{
	ID:         1,
	UserID:     "1",
	MerchantID: "1",
	Amount:     50000,
	CreatedAt:  time.Now().Truncate(time.Second),
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

func (suite *TransactionControllerTestSuite) TestListHandler_Success() {
	mockTransactions := []dto.TransactionDto{mockTransactionDto}
	mockPaging := model.Paging{
		Page:        1,
		RowsPerPage: 10,
		TotalRows:   1,
		TotalPages:  1,
	}

	suite.tum.On("FindAllTransactionByCustomer", 1, 10, "1").Return(mockTransactions, mockPaging, nil)

	suite.amm.On("RequireToken").Return(func(c *gin.Context) {
		c.Set("user", "1")
	})

	suite.cx.Request = httptest.NewRequest(http.MethodGet, "/api/v1/transaction", nil)
	responseRecorder := httptest.NewRecorder()

	suite.tc.listHandler(suite.cx)

	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *TransactionControllerTestSuite) TestListHandler_Fail() {
	mockTransactiondto := []dto.TransactionDto{mockTransactionDto}
	mockError := errors.New("something went wrong")

	suite.tum.On("FindAllTransactionByCustomer", 1, 10, "1").Return(mockTransactiondto, model.Paging{}, mockError)
	suite.amm.On("RequireToken").Return(func(c *gin.Context) {
		c.Set("user", "1")
	})

	suite.cx.Request = httptest.NewRequest(http.MethodGet, "/api/v1/transaction", nil)
	suite.tc.listHandler(suite.cx)
}

func (suite *TransactionControllerTestSuite) TestListHandler_GetUserFail() {
	mockError := errors.New("something went wrong")

	suite.tum.On("FindAllTransactionByCustomer", 1, 10, "1").Return([]dto.TransactionDto{}, model.Paging{}, mockError)
	suite.amm.On("RequireToken").Return(func(c *gin.Context) { c.Set("user", nil) })

	suite.cx.Request = httptest.NewRequest(http.MethodGet, "/api/v1/transaction", nil)
	suite.tc.listHandler(suite.cx)

}

func (suite *TransactionControllerTestSuite) TestCreateHandler_Success() {
	suite.tum.On("RegisterNewTransaction", mockTransaction).Return(mockTransactionDto, nil)

	suite.amm.On("RequireToken").Return(func(c *gin.Context) {})

	requestBody, _ := json.Marshal(mockTransaction)
	suite.cx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/transaction", bytes.NewReader(requestBody))
	responseRecorder := httptest.NewRecorder()

	suite.tc.createHandler(suite.cx)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *TransactionControllerTestSuite) TestCreateHandler_BindFail() {
	suite.tum.On("RegisterNewTransaction", mockTransaction).Return(mockTransactionDto, nil)

	suite.cx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/transaction", nil)

	responseRecorder := httptest.NewRecorder()

	suite.tc.createHandler(suite.cx)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *TransactionControllerTestSuite) TestCreateHandler_Fail() {
	mockError := errors.New("example error message")
	suite.tum.On("RegisterNewTransaction", mockTransaction).Return(mockTransactionDto, mockError)

	suite.amm.On("RequireToken").Return(func(c *gin.Context) {})

	requestBody, _ := json.Marshal(mockTransaction)
	suite.cx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/transaction", bytes.NewReader(requestBody))
	suite.tc.createHandler(suite.cx)
}

func TestTransactionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionControllerTestSuite))
}
