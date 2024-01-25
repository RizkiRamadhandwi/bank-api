package usecase

import (
	"bank-api/entity/dto"
	"bank-api/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Logout(token dto.AuthResponseDto) error
}

type authUseCase struct {
	uc         UserUseCase
	jwtService service.JwtService
}

func (au *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := au.uc.FindUserForLogin(payload.Username, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := au.jwtService.CreateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	return token, nil
}

func (au *authUseCase) Logout(token dto.AuthResponseDto) error {

	err := au.jwtService.DeleteToken(token)
	if err != nil {
		return err
	}
	return nil
}

func NewAuthUseCase(uc UserUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{uc: uc, jwtService: jwtService}
}
