package providers

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/marktrs/vwap-calculation-engine/config"
	"github.com/marktrs/vwap-calculation-engine/domain"
	"github.com/pkg/errors"
)

type Coinbase struct {
	conn *websocket.Conn
	conf config.CoinbaseConfig
}

func NewCoinbase(conf config.CoinbaseConfig) (*Coinbase, error) {
	var wsDialer websocket.Dialer

	conn, _, err := wsDialer.Dial(conf.WebsocketEndpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to establish Coinbase Websocket connection")
	}

	return &Coinbase{
		conn,
		conf,
	}, nil
}

func (c Coinbase) Subscribe(pairs []domain.Pair) error {
	var products []string
	for _, v := range pairs {
		products = append(products, v.String())
	}

	req := coinbaseRequestMessage{
		Type: c.conf.MessageSubscribeType,
		Channels: []coinbaseChannel{
			{
				Name:     c.conf.MatchesChannelName,
				Products: products,
			},
		},
	}

	if err := c.conn.WriteJSON(req); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to subscribe with %v pairs", products))
	}

	return nil
}

func (c Coinbase) Pull(ch chan domain.Trading) error {
	for {
		message := coinbaseResponseMessage{}
		if err := c.conn.ReadJSON(&message); err != nil {
			close(ch)
			return errors.Wrap(err, "unable to parse message")
		}

		if message.Type == c.conf.MessageErrorType {
			close(ch)
			return fmt.Errorf("coinbase error: %v", message)
		}

		if message.Type == c.conf.MessageSubscriptionsType {
			continue
		}

		trading, err := domain.NewTrading(
			message.TradeID,
			message.ProductID,
			message.Size,
			message.Price,
			message.Time,
		)

		if err != nil {
			close(ch)
			return errors.Wrap(err, "unable to parse trading data")
		}

		ch <- trading
	}
}
