package service

import (
	"bank-api/config"
	"bank-api/entity"
	"bank-api/entity/dto"
	"bank-api/shared/model"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	CreateToken(author entity.User) (dto.AuthResponseDto, error)
	ParseToken(tokenHeader string) (jwt.MapClaims, error)
	DeleteToken(token dto.AuthResponseDto) error
	IsTokenRevoked(tokenHeader string) bool
}

type jwtService struct {
	cfg           config.TokenConfig
	revokedTokens map[string]bool
}

func (j *jwtService) CreateToken(user entity.User) (dto.AuthResponseDto, error) {
	claims := model.MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtExpiresTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.ID,
	}

	token := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(j.cfg.JwtSignatureKy)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("oops, failed to create token: %v", err)
	}
	return dto.AuthResponseDto{Token: ss}, nil
}

func (j *jwtService) ParseToken(tokenHeader string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("oops, unexpected signing method: %v", token.Header["alg"])
		}
		return j.cfg.JwtSignatureKy, nil
	})

	if err != nil {
		return nil, fmt.Errorf("oops, failed to verify token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("oops, failed to parse token claims")
	}

	// Periksa apakah token telah kadaluwarsa
	if !token.Valid {
		return nil, fmt.Errorf("oops, token telah kadaluwarsa")
	}

	// Di dalam metode ParseToken
	tokenString := strings.TrimSpace(tokenHeader)

	// Periksa apakah token telah dicabut
	if j.revokedTokens[tokenString] {
		return nil, fmt.Errorf("oops, token telah dicabut")
	}
	return claims, nil
}

func (j *jwtService) DeleteToken(token dto.AuthResponseDto) error {
	j.revokedTokens[token.Token] = true
	return nil
}

func (j *jwtService) IsTokenRevoked(tokenHeader string) bool {
	tokenString := strings.TrimSpace(tokenHeader)
	return j.revokedTokens[tokenString]
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg, revokedTokens: make(map[string]bool)}
}
