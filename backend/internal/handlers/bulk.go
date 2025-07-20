package handlers

import (
	"github.com/bulbasor1509/stock-scoop/backend/internal/crawler"
	"github.com/bulbasor1509/stock-scoop/backend/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BulkHandler(ctx *gin.Context) {
	bulk, err := crawler.BulkCrawler("https://www.nseindia.com/report-detail/display-bulk-and-block-deals")

	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err)
	}
	response.Success(ctx, http.StatusOK, bulk)
}
