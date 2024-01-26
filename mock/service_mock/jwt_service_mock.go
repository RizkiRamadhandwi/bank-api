package service_mock

import (
	"bank-api/entity"
	"bank-api/entity/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type JwtServiceMock struct {
	mock.Mock
}

func (j *JwtServiceMock) CreateToken(author entity.User) (dto.AuthResponseDto, error) {
	args := j.Called(author)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (j *JwtServiceMock) ParseToken(tokenHeader string) (jwt.MapClaims, error) {
	args := j.Called(tokenHeader)
	return args.Get(0).(jwt.MapClaims), args.Error(1)
}

func (j *JwtServiceMock) DeleteToken(token dto.AuthResponseDto) error {
	args := j.Called(token)
	return args.Error(0)
}

func (j *JwtServiceMock) IsTokenRevoked(tokenHeader string) bool {
	args := j.Called(tokenHeader)
	return args.Bool(0)
}
