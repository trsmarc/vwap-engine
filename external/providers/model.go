package providers

import "time"

type coinbaseChannel struct {
	Name     string   `json:"name"`
	Products []string `json:"product_ids"`
}

type coinbaseRequestMessage struct {
	Type     string            `json:"type"`
	Channels []coinbaseChannel `json:"channels"`
}

type coinbaseResponseMessage struct {
	Type      string    `json:"type"`
	TradeID   int       `json:"trade_id"`
	ProductID string    `json:"product_id"`
	Size      string    `json:"size"`
	Price     string    `json:"price"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time"`
}
