package notifiers

import (
	"log"

	"github.com/marktrs/vwap-calculation-engine/domain"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (p Logger) Stream(trading domain.Trading, f float64) error {
	log.Printf("[%s] vwap:%f\n", trading.Pair, f)
	return nil
}
