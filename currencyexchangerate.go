package alphavantage

import (
	"encoding/json"
	"fmt"
)

// InsiderTransaction represents a single insider transaction record.
type CurrencyExchangeRate struct {
	FromCurrencyCode string `json:"1. From_Currency Code"`
	FromCurrencyName string `json:"2. From_Currency Name"`
	ToCurrencyCode   string `json:"3. To_Currency Code"`
	ToCurrencyName   string `json:"4. To_Currency Name"`
	ExchangeRate     string `json:"5. Exchange Rate"`
	LastRefreshed    string `json:"6. Last Refreshed"`
	TimeZone         string `json:"7. Time Zone"`
	BidPrice         string `json:"8. Bid Price"`
	AskPrice         string `json:"9. Ask Price"`
}

type currencyExchangeRateResponse struct {
	Rate CurrencyExchangeRate `json:"Realtime Currency Exchange Rate"`
}

func toCurrencyExchangeRate(buf []byte) (*CurrencyExchangeRate, error) {
	var currencyExchangeRateResponse currencyExchangeRateResponse
	if err := json.Unmarshal(buf, &currencyExchangeRateResponse); err != nil {
		return nil, err
	}
	return &currencyExchangeRateResponse.Rate, nil
}

// InsiderTransactions fetches and returns the insider transactions data for the specified symbol.
func (c *Client) CurrencyExchangeRate(fromCurrency string, toCurrency string) (*CurrencyExchangeRate, error) {
	const function = "CURRENCY_EXCHANGE_RATE"
	url := fmt.Sprintf("%s/query?function=%s&from_currency=%s&to_currency=%s&apikey=%s", baseURL, function, fromCurrency, toCurrency, c.apiKey)

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toCurrencyExchangeRate(body)
}
