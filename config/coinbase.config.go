package config

type CoinbaseConfig struct {
	WebsocketEndpoint        string
	MatchesChannelName       string
	MessageSubscribeType     string
	MessageSubscriptionsType string
	MessageErrorType         string
}

var Coinbase = CoinbaseConfig{
	WebsocketEndpoint:        "wss://ws-feed.pro.coinbase.com",
	MatchesChannelName:       "matches",
	MessageSubscribeType:     "subscribe",
	MessageSubscriptionsType: "subscriptions",
	MessageErrorType:         "error",
}
