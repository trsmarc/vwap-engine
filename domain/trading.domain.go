package domain

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Trading struct {
	ID        int
	Pair      string
	Share     float64
	Price     float64
	CreatedAt time.Time
}

func NewTrading(tradeID int, pair string, size string, price string, createdAt time.Time) (Trading, error) {
	shareNumber, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "unable to convert size to float64")
	}

	priceNumber, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "unable to convert price to float64")
	}

	return Trading{ID: tradeID, Pair: pair, Share: shareNumber, Price: priceNumber, CreatedAt: createdAt}, nil
}
