package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessRes struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SetSuccessResponse(ctx *gin.Context, httpStatus int, data interface{}) {
	// Prepare Response
	response := SuccessRes{
		Code:   httpStatus,
		Status: http.StatusText(httpStatus),
		Data:   data,
	}

	ctx.JSON(response.Code, response)
}
