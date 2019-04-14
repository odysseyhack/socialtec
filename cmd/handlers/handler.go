package handlers

import (
	"github.com/odysseyhack/socialtec/node/pkg/market"
	"github.com/odysseyhack/socialtec/node/pkg/store"
)

type Handler struct {
	market market.Market
	store  store.Store
}

func NewHandler(market market.Market, store store.Store) *Handler {
	return &Handler{market: market, store: store}
}
