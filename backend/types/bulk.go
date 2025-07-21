package types

import (
	"github.com/bulbasor1509/stock-scoop/backend/internal/database"
	"github.com/google/uuid"
	"time"
)

type Stock struct {
	Date     string `json:"date"`
	Symbol   string `json:"symbol"`
	Security string `json:"security"`
	Client   string `json:"client"`
	Type     string `json:"type"`
	TradeQty string `json:"trade_qty"`
	AvgPrice string `json:"avg_price"`
	Remark   string `json:"remark"`
}

type StocksInsert struct {
	Dates     []time.Time
	Symbols   []string
	Security  []string
	Clients   []string
	Types     []database.Trade
	TradeQtys []string
	AvgPrices []string
	Remarks   []string
}

type StockResponse struct {
	Id       uuid.UUID      `json:"id"`
	Date     time.Time      `json:"date"`
	Symbol   string         `json:"symbol"`
	Security string         `json:"security"`
	Client   string         `json:"client"`
	Type     database.Trade `json:"type"`
	TradeQty string         `json:"trade_qty"`
	AvgPrice string         `json:"avg_price"`
}

func BulkToJSON(bulk database.BulkDeal) StockResponse {
	return StockResponse{
		Id:       bulk.ID,
		Date:     bulk.Date,
		Symbol:   bulk.Symbol,
		Security: bulk.SecurityName,
		Client:   bulk.ClientName,
		Type:     bulk.TradeType,
		TradeQty: bulk.QuantityTraded,
		AvgPrice: bulk.WtPrice,
	}
}

func DatabaseBulkToBulk(bulks []database.BulkDeal) []StockResponse {
	var response []StockResponse
	for _, bulk := range bulks {
		response = append(response, BulkToJSON(bulk))
	}
	return response
}
