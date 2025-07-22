package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRequestParser(url string, target interface{}) error {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Referer", "https://www.nseindia.com/")

	resp, err := client.Do(request)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	return nil
}
