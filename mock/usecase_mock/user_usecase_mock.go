package usecase_mock

import (
	"bank-api/entity"

	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (u *UserUsecaseMock) FindUserForLogin(username, password string) (entity.User, error) {
	args := u.Called(username, password)
	return args.Get(0).(entity.User), args.Error(1)
}
