package types

type BulkDeal struct {
	Date       string  `json:"BD_DT_DATE"`
	OrderTime  string  `json:"BD_DT_ORDER"`
	Symbol     string  `json:"BD_SYMBOL"`
	ScripName  string  `json:"BD_SCRIP_NAME"`
	ClientName string  `json:"BD_CLIENT_NAME"`
	BuySell    string  `json:"BD_BUY_SELL"`
	QtyTraded  int     `json:"BD_QTY_TRD"`
	WtPrice    float64 `json:"BD_TP_WATP"`
	Remarks    string  `json:"BD_REMARKS"`
}

type BulkDealResponse struct {
	Data []BulkDeal `json:"data"`
}
