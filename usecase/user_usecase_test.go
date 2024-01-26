package usecase

import (
	"bank-api/entity"
	"bank-api/mock/repository_mock"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	urm *repository_mock.UserRepoMock
	uuc UserUseCase
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.urm = new(repository_mock.UserRepoMock)
	suite.uuc = NewUserUseCase(suite.urm)
}

var mockUser = entity.User{
	ID:       "1",
	Name:     "Rizki",
	Username: "Dwi",
	Password: "Ramadhan",
}

func (suite *UserUsecaseTestSuite) TestFindUserForLogin_Success() {
	suite.urm.On("GetForLogin", mockUser.Username, mockUser.Password).Return(mockUser, nil)

	actual, err := suite.uuc.FindUserForLogin(mockUser.Username, mockUser.Password)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockUser, actual)
}

func (suite *UserUsecaseTestSuite) TestFindUserForLogin_NoFieldFail() {
	suite.urm.On("GetForLogin", "", "").Return(entity.User{}, errors.New("username and password required"))

	_, err := suite.uuc.FindUserForLogin("", "")
	assert.Error(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestFindUserForLogin_NoFieldUsernameFail() {
	suite.urm.On("GetForLogin", "", mockUser.Password).Return(entity.User{}, errors.New("username required"))

	_, err := suite.uuc.FindUserForLogin("", mockUser.Password)
	assert.Error(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestFindUserForLogin_NoFieldPasswordFail() {
	suite.urm.On("GetForLogin", mockUser.Username, "").Return(entity.User{}, errors.New("password required"))

	_, err := suite.uuc.FindUserForLogin(mockUser.Username, "")
	assert.Error(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestFindUserForLogin_NoExistsFail() {
	suite.urm.On("GetForLogin", "rizki", "dwi").Return(entity.User{}, fmt.Errorf("user doesn't exists"))

	_, err := suite.uuc.FindUserForLogin("rizki", "dwi")
	assert.Error(suite.T(), err)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
