package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/bulbasor1509/stock-scoop/backend/types"
	"strconv"
	"strings"
	"time"
)

func TableParser(table string) ([]types.Stock, error) {
	var bulk []types.Stock

	document, err := goquery.NewDocumentFromReader(strings.NewReader(table))
	if err != nil {
		return nil, err
	}
	document.Find("tr").Each(func(i int, row *goquery.Selection) {
		var rowData []string
		if i == 0 {
			return
		}

		row.Find("td").Each(func(j int, s *goquery.Selection) {
			text := strings.TrimSpace(s.Text())
			rowData = append(rowData, text)
		})

		if len(rowData) < 7 {
			return
		}

		stock := types.Stock{
			Date:     rowData[0],
			Symbol:   rowData[1],
			Security: rowData[2],
			Client:   rowData[3],
			Type:     rowData[4],
			TradeQty: rowData[5],
			AvgPrice: rowData[6],
		}
		bulk = append(bulk, stock)
	})
	return bulk, nil
}

func ParseIndianCurrency(currency string) (int64, error) {
	cur := strings.ReplaceAll(currency, ",", "")
	return strconv.ParseInt(cur, 10, 64)
}

func DateParser(date string) (time.Time, error) {
	layout := "02-Jan-2006"
	parts := strings.Split(date, "-")

	parts[1] = strings.ToUpper(parts[1][:1]) + strings.ToLower(parts[1][1:])
	formatted := strings.Join(parts, "-")

	return time.Parse(layout, formatted)
}
