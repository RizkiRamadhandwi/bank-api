package repository_mock

import (
	"bank-api/entity"
	"bank-api/entity/dto"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) GetForLogin(username, password string) (entity.User, error) {
	args := u.Called(username, password)
	return args.Get(0).(entity.User), args.Error(1)
}
func (u *UserRepoMock) GetByIdCust(id string) (dto.UserDto, error) {
	args := u.Called(id)
	return args.Get(0).(dto.UserDto), args.Error(1)
}
