package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorRes struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

func SetErrorResponse(ctx *gin.Context, httpStatus int, Error string) {
	// Prepare Response
	response := ErrorRes{
		Code:   httpStatus,
		Status: http.StatusText(httpStatus),
		Error:  Error,
	}

	ctx.JSON(response.Code, response)
}
