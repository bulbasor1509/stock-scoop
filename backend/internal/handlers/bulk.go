package handlers

import (
	"github.com/bulbasor1509/stock-scoop/backend/internal/crawler"
	"github.com/bulbasor1509/stock-scoop/backend/internal/helpers"
	"github.com/bulbasor1509/stock-scoop/backend/internal/response"
	"github.com/bulbasor1509/stock-scoop/backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BulkHandler(cfg *types.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bulk, err := crawler.BulkCrawler("https://www.nseindia.com/report-detail/display-bulk-and-block-deals")

		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err)
			return
		}

		if err := helpers.InsertBulkData(ctx, cfg, bulk); err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err)
			return
		}

		response.Success(ctx, http.StatusOK, bulk)
	}
}

func GetBulkHandler(cfg *types.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bulk, err := cfg.DB.GetBulkData(ctx)
		bulkRes := types.DatabaseBulkToBulk(bulk)

		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err)
		}
		response.Success(ctx, http.StatusOK, bulkRes)
	}
}
