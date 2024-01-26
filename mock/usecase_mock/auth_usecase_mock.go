package usecase_mock

import (
	"bank-api/entity/dto"

	"github.com/stretchr/testify/mock"
)

type AuthUsecaseMock struct {
	mock.Mock
}

func (a *AuthUsecaseMock) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	args := a.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (a *AuthUsecaseMock) Logout(token dto.AuthResponseDto) error {
	args := a.Called(token)
	return args.Error(0)
}
