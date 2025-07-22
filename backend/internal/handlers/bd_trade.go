package handlers

import (
	"fmt"
	"github.com/bulbasor1509/stock-scoop/backend/internal/helpers"
	"github.com/bulbasor1509/stock-scoop/backend/internal/response"
	"github.com/bulbasor1509/stock-scoop/backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	BASE_URL = "https://www.nseindia.com/api"
)

func GetBDDataHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var result types.BulkDealResponse

		today := time.Now().Truncate(24 * time.Hour)
		yesterday := today.AddDate(0, 0, -1)
		format := "02-01-2006"

		bd_url := fmt.Sprintf("%v/historicalOR/bulk-block-short-deals?optionType=bulk_deals&from=%v&to=%v", BASE_URL, yesterday.Format(format), today.Format(format))
		err := helpers.GetRequestParser(bd_url, &result)
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err)
			return
		}
		response.Success(ctx, http.StatusOK, result.Data)
	}
}
