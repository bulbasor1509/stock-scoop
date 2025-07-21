package helpers

import (
	"github.com/bulbasor1509/stock-scoop/backend/internal/database"
	"github.com/bulbasor1509/stock-scoop/backend/internal/parser"
	"github.com/bulbasor1509/stock-scoop/backend/types"
	"github.com/gin-gonic/gin"
)

const (
	layout = "02-JAN-2006"
)

func FormatBulkData(stocks []types.Stock) types.StocksInsert {
	var stocksData types.StocksInsert
	for _, stock := range stocks {
		parsedDate, _ := parser.DateParser(stock.Date)
		stocksData.Dates = append(stocksData.Dates, parsedDate)
		stocksData.Symbols = append(stocksData.Symbols, stock.Symbol)
		stocksData.Security = append(stocksData.Security, stock.Security)
		stocksData.Clients = append(stocksData.Clients, stock.Client)

		if stock.Type == "BUY" {
			stocksData.Types = append(stocksData.Types, database.TradeBUY)
		} else {
			stocksData.Types = append(stocksData.Types, database.TradeSELL)
		}

		stocksData.TradeQtys = append(stocksData.TradeQtys, stock.TradeQty)
		stocksData.AvgPrices = append(stocksData.AvgPrices, stock.AvgPrice)
		stocksData.Remarks = append(stocksData.Remarks, stock.Remark)
	}

	return stocksData
}

func InsertBulkData(ctx *gin.Context, cfg *types.ApiConfig, stocks []types.Stock) error {
	stocksData := FormatBulkData(stocks)
	err := cfg.DB.InsertBulkData(ctx, database.InsertBulkDataParams{
		Dates:         stocksData.Dates,
		Symbols:       stocksData.Symbols,
		SecurityNames: stocksData.Security,
		ClientNames:   stocksData.Clients,
		TradeTypes:    stocksData.Types,
		Quantities:    stocksData.TradeQtys,
		Prices:        stocksData.AvgPrices,
	})
	if err != nil {
		return err
	}
	return nil
}
