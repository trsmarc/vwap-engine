package repository

import (
	"github.com/marktrs/vwap-calculation-engine/domain"
)

type Notifier interface {
	Stream(domain.Trading, float64) error
}
