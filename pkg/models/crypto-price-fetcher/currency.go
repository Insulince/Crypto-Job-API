package crypto_price_fetcher

type Currency struct {
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	ID                string `json:"id"`
	CurrentPriceInUSD string `json:"current-price-in-usd"`
	CurrentPriceInBTC string `json:"current-price-in-btc"`
}
