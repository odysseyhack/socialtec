package handlers

import (
	"net/http"

	"github.com/odysseyhack/socialtec/node/pkg/store"

	"github.com/go-chi/render"
	"github.com/odysseyhack/socialtec/node/pkg/market"
)

func (handler Handler) MyOffers(w http.ResponseWriter, r *http.Request) {
	var offers []market.Offer
	if err := handler.store.Get("my_offers", &offers); err != nil && err != store.ErrorNotFound {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	render.JSON(w, r, offers)
}

func (handler Handler) MyInterests(w http.ResponseWriter, r *http.Request) {
	var offers []market.Offer
	if err := handler.store.Get("my_interests", &offers); err != nil && err != store.ErrorNotFound {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	render.JSON(w, r, offers)
}

func (handler Handler) addMyOffer(offer market.Offer) error {
	var offers []market.Offer
	if err := handler.store.Get("my_offers", &offers); err != nil && err != store.ErrorNotFound {
		return err
	}
	offers = append(offers, offer)
	return handler.store.Set("my_offers", offers)
}

func (handler Handler) addMyIntrest(offer market.Offer) error {
	var offers []market.Offer
	if err := handler.store.Get("my_interests", &offers); err != nil && err != store.ErrorNotFound {
		return err
	}
	offers = append(offers, offer)
	return handler.store.Set("my_interests", offers)
}
