package handlers

import (
	"github.com/odysseyhack/socialtec/node/pkg/market"
)

type Handler struct {
	market market.Market
}

func NewHandler(market market.Market) *Handler {
	return &Handler{market:market}
}