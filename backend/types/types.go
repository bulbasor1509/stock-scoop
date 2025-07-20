package types

type Stock struct {
	Date     string `json:"date"`
	Symbol   string `json:"symbol"`
	Security string `json:"security"`
	Client   string `json:"client"`
	Type     string `json:"type"`
	TradeQty string `json:"trade_qty"`
	AvgPrice string `json:"avg_price"`
	Remarks  string `json:"remarks"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
