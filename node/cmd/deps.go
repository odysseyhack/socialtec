package main

import (
	"fmt"

	"github.com/odysseyhack/socialtec/node/pkg/market"
	"github.com/odysseyhack/socialtec/node/pkg/store"
)

func initDependencies(conf *Config) {
	market.DefaultMarket = market.NewMarket(
		"ws://0.0.0.0:7545",
		conf.NodeKey,
		// "3cae6c4f79bd2c11028cba295938d92bef1ccc9691b95b98f7759df3b855a4cc",
		//	"536d2fffff9af2dcb66e75782ccf75450246703130b8ab775f1f5893a6cef26a",
		"0x40B1Ba95fa7144c4315D0C6e97F4b29C7989feE6",
	)
	// market.DefaultMarket = market.NewMarket(
	// 	"wss://rinkeby.infura.io/ws",
	// 	"d3193a12ca4dd3854639cf0524ac4a06083f703e63c5d55f6c20624a27bb29aa",
	// 	"0xC7585A3569202e8Df12E3323AB280ABB2FA4E59a",
	// )

	// market.DefaultMarket = market.NewMarket(
	// 	"http://172.16.166.127:7545",
	// 	"8678dc180d0a126d14e861a1afe9750071f0adfa495ac3f06355680595a65283",
	// 	"0xB91fE4dd0c31CA4366c7e4bD8d723c1d278BAE66",
	// )
	store.Default = store.NewStore(fmt.Sprintf("./%s.db", conf.NodeKey[:10]))
	go market.DefaultMarket.ListenBroadcasts(store.Default)
}
