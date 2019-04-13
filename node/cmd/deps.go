package main

import "github.com/odysseyhack/socialtec/node/pkg/market"

func initDependencies(conf *Config) {
	market.DefaultMarket = market.NewMarket(
		"ws://0.0.0.0:7545",
		"536d2fffff9af2dcb66e75782ccf75450246703130b8ab775f1f5893a6cef26a",
	)
}
