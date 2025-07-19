package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"net/http"
	"strings"
)

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

var bulk []Stock

func tableParser(table string) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(table))
	if err != nil {
		panic(err)
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
		stock := Stock{
			Date:     rowData[0],
			Symbol:   rowData[1],
			Security: rowData[2],
			Client:   rowData[3],
			Type:     rowData[4],
			TradeQty: rowData[5],
			Remarks:  rowData[6],
		}
		bulk = append(bulk, stock)
	})
}

func open_browser(url string) {
	browser := rod.New().MustConnect()
	defer browser.Close()

	page := browser.MustPage(url).MustWaitLoad()
	page.MustElement("a[data-val=\"1D\"]").MustClick()
	rowsHtml, _ := page.MustElement("#HistBulkBlockDataTable").MustWaitVisible().HTML()
	tableParser(rowsHtml)
}

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		open_browser("https://www.nseindia.com/report-detail/display-bulk-and-block-deals")
		c.JSON(http.StatusOK, gin.H{
			"bulk": bulk,
		})
	})

	if err := router.Run(":8080"); err != nil {
		return
	}
}
