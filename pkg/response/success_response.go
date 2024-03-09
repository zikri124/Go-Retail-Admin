package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResTemplate struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SetSuccessResponse(ctx *gin.Context, httpStatus int, data interface{}) {
	// Prepare Response
	response := SuccessResTemplate{
		Code:   httpStatus,
		Status: http.StatusText(httpStatus),
		Data:   data,
	}

	ctx.JSON(response.Code, response)
}
