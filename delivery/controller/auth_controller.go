package controller

import (
	"bank-api/config"
	"bank-api/entity/dto"
	"bank-api/logging"
	"bank-api/shared/common"
	"bank-api/shared/service"
	"bank-api/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUc     usecase.AuthUseCase
	rg         *gin.RouterGroup
	jwtService service.JwtService
}

func (ac *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := ac.authUc.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, token, "Ok")

	activity := "logging in"
	logging.LogUserAuth(activity)
}

func (ac *AuthController) logoutHandler(ctx *gin.Context) {
	tokenHeader := ctx.GetHeader("Authorization")
	if tokenHeader == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Token required")
		return
	}

	token := strings.Replace(tokenHeader, "Bearer ", "", 1)
	_, err := ac.jwtService.ParseToken(token)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
		return
	}

	err = ac.authUc.Logout(dto.AuthResponseDto{Token: token})
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = ac.jwtService.ParseToken(token)
	if err == nil {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, "Token still valid after logout")
		return
	}

	common.SendNoContentResponse(ctx)

	activity := "logging out"
	logging.LogUserAuth(activity)
}

func (a *AuthController) Route() {
	a.rg.POST(config.PostLogin, a.loginHandler)
	a.rg.POST(config.PostLogout, a.logoutHandler)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup, jwtService service.JwtService) *AuthController {
	return &AuthController{authUc: authUc, rg: rg, jwtService: jwtService}
}
