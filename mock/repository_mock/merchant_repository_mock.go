package repository_mock

import (
	"bank-api/entity/dto"

	"github.com/stretchr/testify/mock"
)

type MerchantRepoMock struct {
	mock.Mock
}

func (m *MerchantRepoMock) GetByIdMerc(id string) (dto.MerchantDto, error) {
	args := m.Called(id)
	return args.Get(0).(dto.MerchantDto), args.Error(1)
}
