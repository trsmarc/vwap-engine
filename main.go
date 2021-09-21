package main

import (
	"log"

	"github.com/marktrs/vwap-calculation-engine/config"
	"github.com/marktrs/vwap-calculation-engine/domain"
	"github.com/marktrs/vwap-calculation-engine/external/notifiers"
	"github.com/marktrs/vwap-calculation-engine/external/providers"
	"github.com/marktrs/vwap-calculation-engine/service"
)

func main() {
	provider, err := providers.NewCoinbase(config.Coinbase)
	if err != nil {
		log.Fatalln(err)
	}

	if err = provider.Subscribe([]domain.Pair{
		domain.NewPair("BTC", "USD"),
		domain.NewPair("ETH", "USD"),
		domain.NewPair("ETH", "BTC"),
	}); err != nil {
		log.Fatalln(err)
	}

	logger := notifiers.NewLogger()
	vwapService := service.NewVWAPService(provider, logger, config.App)
	vwapService.Calculate()
}
