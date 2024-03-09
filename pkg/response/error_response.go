package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResTemplate struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

func SetErrorResponse(ctx *gin.Context, httpStatus int, err error) {
	// Prepare Response
	response := ErrorResTemplate{
		Code:   httpStatus,
		Status: http.StatusText(httpStatus),
		Error:  err.Error(),
	}

	ctx.JSON(response.Code, response)
}
