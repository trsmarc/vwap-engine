package tests

import "github.com/marktrs/vwap-calculation-engine/domain"

type notifierMock struct {
	ch chan float64
}

func (n notifierMock) Stream(_ domain.Trading, f float64) error {
	n.ch <- f

	return nil
}
