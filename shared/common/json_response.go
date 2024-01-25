package common

import (
	"bank-api/shared/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCreateResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusCreated,
			Message: "Created",
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendPagedResponse(ctx *gin.Context, data []interface{}, paging model.Paging, message string) {
	ctx.JSON(http.StatusOK, &model.PagedResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}

func SendNoContentResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, &model.Status{
		Code:    http.StatusNoContent,
		Message: "No Content",
	})
}
