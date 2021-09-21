package repository

import (
	"github.com/marktrs/vwap-calculation-engine/domain"
)

type Provider interface {
	Pull(chan domain.Trading) error
	Subscribe([]domain.Pair) error
}
