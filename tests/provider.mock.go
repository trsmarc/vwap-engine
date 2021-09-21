package tests

import "github.com/marktrs/vwap-calculation-engine/domain"

type providerMock struct {
	tradings []domain.Trading
}

func (p providerMock) Subscribe(_ []domain.Pair) error {
	return nil
}


func (p providerMock) Pull(ch chan domain.Trading) error {
	if p.tradings == nil {
		return nil
	}

	for _, v := range p.tradings {
		ch <- v
	}

	return nil
}
