package service

import (
	"log"

	"github.com/marktrs/vwap-calculation-engine/config"
	"github.com/marktrs/vwap-calculation-engine/domain"
	repo "github.com/marktrs/vwap-calculation-engine/repository"
)

type VWAPService struct {
	provider repo.Provider
	notifier repo.Notifier
	conf     config.AppConfig
}

func NewVWAPService(provider repo.Provider, notifier repo.Notifier, conf config.AppConfig) *VWAPService {
	return &VWAPService{
		provider: provider,
		notifier: notifier,
		conf:     conf,
	}
}

func (s *VWAPService) Calculate() {
	var (
		ch          = make(chan domain.Trading)
		operations  = map[string]*domain.Queue{}
		volume      = map[string]float64{}
		totalVolume = map[string]float64{}
	)

	go func() {
		if err := s.provider.Pull(ch); err != nil {
			log.Fatalln(err)
		}
	}()

	for trading := range ch {
		if _, ok := operations[trading.Pair]; !ok {
			operations[trading.Pair] = domain.NewQueue()
		}

		if _, ok := volume[trading.Pair]; !ok {
			volume[trading.Pair] = 0
		}

		if _, ok := totalVolume[trading.Pair]; !ok {
			totalVolume[trading.Pair] = 0
		}

		operations[trading.Pair].Add(trading)

		volume[trading.Pair] += trading.Price * trading.Share
		totalVolume[trading.Pair] += trading.Share

		if (s.conf.MaxWindowSize > 0) && (operations[trading.Pair].Len() > s.conf.MaxWindowSize) {
			var first = operations[trading.Pair].Remove()

			volume[trading.Pair] -= first.Price * first.Share
			totalVolume[trading.Pair] -= first.Share
		}

		var result = volume[trading.Pair] / totalVolume[trading.Pair]

		if err := s.notifier.Stream(trading, result); err != nil {
			log.Println("unable to stream calculation result:", err.Error())
		}
	}
}
