package controller

import (
	"bank-api/config"
	"bank-api/delivery/middleware"
	"bank-api/entity"
	"bank-api/logging"
	"bank-api/shared/common"
	"bank-api/usecase"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transUc usecase.TransactionUseCase
	rg      *gin.RouterGroup
	authMid middleware.AuthMiddleware
}

func (tc *TransactionController) createHandler(ctx *gin.Context) {

	var payload entity.Transaction
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := ctx.MustGet("user").(string)
	payload.UserID = user
	rsv, err := tc.transUc.RegisterNewTransaction(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, rsv)

	activity := "has already made a transaction"
	logging.LogUserActivity(user, activity)
}

func (tc *TransactionController) listHandler(ctx *gin.Context) {
	userValue, _ := ctx.Get("user")
	user, _ := userValue.(string)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

	rsv, paging, err := tc.transUc.FindAllTransactionByCustomer(page, size, user)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var interfaceSlice = make([]interface{}, len(rsv))
	for i, v := range rsv {
		interfaceSlice[i] = v
	}

	common.SendPagedResponse(ctx, interfaceSlice, paging, "Ok")

	activity := "is viewing the list transaction"
	logging.LogUserActivity(user, activity)

}

func (tc *TransactionController) Route() {
	tc.rg.Use(tc.authMid.RequireToken())
	tc.rg.POST(config.PostTrasaction, tc.createHandler)
	tc.rg.GET(config.GetTransaction, tc.listHandler)
}

func NewTransactionController(transUc usecase.TransactionUseCase, rg *gin.RouterGroup, authMid middleware.AuthMiddleware) *TransactionController {
	return &TransactionController{transUc: transUc, rg: rg, authMid: authMid}
}
