package tests

import (
	"fmt"
	"testing"

	"github.com/marktrs/vwap-calculation-engine/config"
	"github.com/marktrs/vwap-calculation-engine/domain"
	"github.com/marktrs/vwap-calculation-engine/service"
)

func TestCalculation(t *testing.T) {
	var (
		btcUSD = fmt.Sprintf("%s-%s", "BTC", "USD")
		ethUSD = fmt.Sprintf("%s-%s", "ETH", "USD")

		tradingsMock = []domain.Trading{
			{Share: 1, Price: 2, Pair: btcUSD},
			{Share: 1, Price: 2, Pair: btcUSD},
			{Share: 1, Price: 2, Pair: btcUSD},
			{Share: 1, Price: 2, Pair: ethUSD},
		}

		expected = []float64{2, 2, 2, 2}
		results  = make(chan float64)
		i        = 0
	)

	var service = service.NewVWAPService(
		providerMock{tradings: tradingsMock},
		notifierMock{ch: results},
		config.App,
	)

	go service.Calculate()

	for got := range results {
		if got != expected[i] {
			t.Errorf("result = %v, expected = %v", got, expected[i])
		}

		i++

		if len(expected) == i {
			close(results)
		}
	}
}
