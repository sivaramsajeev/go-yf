package yf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"text/template"
	"time"
)

const (
	yahooFinanceAPI = "https://query1.finance.yahoo.com/v8/finance/chart/{{.Symbol}}?range={{.Range}}&includePrePost=false&interval={{.Interval}}&corsDomain=finance.yahoo.com&.tsrc=finance"
)

var (
	// DebugLogging enables verbose output from the yf package when set to true.
	DebugLogging = false
)

// GetStock returns stock details for a particular symbol from the Yahoo Finance API.
func GetStock(symbol, dateRange, interval string) (*Stock, error) {
	resp, err := doRequest(symbol, dateRange, interval)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected StatusCode: %d", resp.StatusCode)
	}

	var target struct {
		Chart struct {
			Result []Stock
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, err
	}

	if len(target.Chart.Result) != 1 {
		return nil, fmt.Errorf("Unexpected response, %d results returned, API change must have occurred", len(target.Chart.Result))
	}

	return &target.Chart.Result[0], nil
}

func doRequest(symbol, dateRange, interval string) (*http.Response, error) {
	tmpl, err := template.New("YF-API").Parse(yahooFinanceAPI)
	if err != nil {
		return nil, err
	}

	p := struct {
		Symbol   string
		Range    string
		Interval string
	}{
		Symbol:   symbol,
		Range:    dateRange,
		Interval: interval,
	}
	var url bytes.Buffer
	if err := tmpl.Execute(&url, p); err != nil {
		return nil, err
	}

	debug(url.String())
	return newReq(url.String())
}

func debug(str string) {
	if !DebugLogging {
		return
	}

	fmt.Println(str)
}

func newReq(urlStr string) (*http.Response, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "MyCustomAgent/1.0 (Go HTTP Client)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}
