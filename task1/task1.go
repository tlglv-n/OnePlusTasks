package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	apiURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
)

type CryptoCurrency struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

func fetchCryptoCurrencies() ([]CryptoCurrency, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var currencies []CryptoCurrency
	if err := json.NewDecoder(resp.Body).Decode(&currencies); err != nil {
		return nil, err
	}

	return currencies, nil
}

func findCurrency(currencies []CryptoCurrency, symbol string) (CryptoCurrency, bool) {
	for _, currency := range currencies {
		if currency.Symbol == symbol {
			return currency, true
		}
	}
	return CryptoCurrency{}, false
}

func updateInterval(done chan struct{}) {
	go func() {
		time.Sleep(10 * time.Minute)
		fmt.Println("Курс обновлен")
		updateInterval(done)
	}()
}

func main() {
	// Создаем канал для сигнала обновления курса каждые 10 минут
	done := make(chan struct{})

	// Запускаем горутину для подсчета и обновления времени
	updateInterval(done)

	for {
		currencies, err := fetchCryptoCurrencies()
		if err != nil {
			fmt.Printf("Error fetching crypto currencies: %v\n", err)
			continue
		}

		// Получаем курс для определенной криптовалюты
		var symbol string
		fmt.Print("Введите символ криптовалюты (например, btc): ")
		_, err = fmt.Scan(&symbol)
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}
		symbol = strings.TrimSpace(symbol)

		currency, found := findCurrency(currencies, symbol)
		if !found {
			fmt.Printf("Currency with symbol %s not found\n", symbol)
			continue
		}

		fmt.Printf("%s (%s): $%.2f\n", currency.Name, currency.Symbol, currency.CurrentPrice)

	}
	close(done)
}
