package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"time"
)

func main() {
	browser := rod.New().MustConnect()
	defer func(browser *rod.Browser) {
		err := browser.Close()
		if err != nil {
			fmt.Printf("Browser close error: %v\n", err)
		}
	}(browser)

	page := browser.MustPage("https://www.nseindia.com/reports-indices-historical-index-data")
	page.MustWaitLoad()
	time.Sleep(1 * time.Second)

	options := page.MustElement("#hpReportIndexTypeSearchInput")
	_ = options.Select([]string{`[value="NIFTY 50"]`}, false, rod.SelectorTypeCSSSector)
	time.Sleep(10 * time.Second)
}
