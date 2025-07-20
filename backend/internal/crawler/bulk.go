package crawler

import (
	"github.com/bulbasor1509/stock-scoop/backend/internal/parser"
	"github.com/bulbasor1509/stock-scoop/backend/types"
	"github.com/go-rod/rod"
	"log"
)

func BulkCrawler(url string) ([]types.Stock, error) {
	browser := rod.New().MustConnect()
	defer func(browser *rod.Browser) {
		err := browser.Close()
		if err != nil {
			log.Printf("browser close error: %s", err)
		}
	}(browser)

	page := browser.MustPage(url).MustWaitLoad()
	page.MustElement("a[data-val=\"1D\"]").MustClick()
	htmlRows, _ := page.MustElement("#HistBulkBlockDataTable").MustWaitVisible().HTML()

	bulk, err := parser.TableParser(htmlRows)
	if err != nil {
		return nil, err
	}

	return bulk, nil
}
