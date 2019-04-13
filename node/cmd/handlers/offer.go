package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/odysseyhack/socialtec/node/pkg/market"
)

func getAllOffers() []market.Offer {
	var offers []market.Offer
	for i := 0; i < 10; i++ {
		offers = append(offers, market.Offer{
			Name:     "dsasas",
			Details:  "sasas" + fmt.Sprint(i),
			ImageURL: "https://via.placeholder.com/150",
			ID:       fmt.Sprint(i),
		})
	}
	return offers
}

func (handler Handler) NewOffer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, getAllOffers()[0])
}

func (handler Handler) GetOffers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, getAllOffers())
}

func (handler Handler) DeleteOffer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]interface{}{"success": true})
}
