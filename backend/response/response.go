package response

import (
	"github.com/bulbasor1509/stock-scoop/backend/types"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, statusCode int, data interface{}) {
	response := types.Response{
		Status:  1,
		Message: "data send successfully",
		Data:    data,
	}
	ctx.JSON(statusCode, &response)
}

func Fail(ctx *gin.Context, statusCode int, err error) {
	response := types.Response{
		Status:  1,
		Message: err.Error(),
		Data:    nil,
	}
	ctx.JSON(statusCode, &response)
}
