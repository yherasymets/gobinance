package main

import (
	"context"
	"fmt"

	"github.com/aiviaio/go-binance/v2"
	"github.com/aiviaio/go-binance/v2/futures"
)

func main() {
	// Initialize Binance Futures client.
	client := binance.NewFuturesClient("", "")
	logger := client.Logger
	ctx := context.Background()

	// Retrieve exchange information.
	exchangeInfo, err := client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		logger.Fatalf("can't retrive exchange info : %v", err)
	}

	// Extract first five symbols.
	symbols := make([]string, 0, 5)
	for _, symbol := range exchangeInfo.Symbols[:5] {
		symbols = append(symbols, symbol.Symbol)
	}

	ch := make(chan *futures.SymbolPrice, 5)
	defer close(ch)

	// Create goroutines to fetch prices for each symbol concurrently.
	for _, symbol := range symbols {

		go func(symbol string) {
			price, err := client.NewListPricesService().Symbol(symbol).Do(ctx)
			if err != nil {
				logger.Fatalf("can't retrive price for %s: %s\n", symbol, err.Error())
			}
			ch <- price[0]
		}(symbol)
	}

	// Final result output.
	for i := 0; i < len(symbols); i++ {
		price := <-ch
		fmt.Println(price.Symbol, price.Price)
	}
}
