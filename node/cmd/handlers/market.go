package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
		})
	}
	return offers
}

func (handler Handler) NewOffer(w http.ResponseWriter, r *http.Request) {
	var offer market.Offer
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	if err := handler.market.AddOffer(offer); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	render.JSON(w, r, offer)
}

func (handler Handler) GetOffers(w http.ResponseWriter, r *http.Request) {
	offers, err := handler.market.GetAvailableOffers(0)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	render.JSON(w, r, offers)
}

func (handler Handler) DeleteOffer(w http.ResponseWriter, r *http.Request) {
	offerID := getOfferID(r)

	err := handler.market.DeleteOffer(offerID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	render.JSON(w, r, map[string]interface{}{"success": true})
}

func (handler Handler) AddInterest(w http.ResponseWriter, r *http.Request) {
	offerID := getOfferID(r)

	err := handler.market.ShowIntrest(offerID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	render.JSON(w, r, map[string]interface{}{"success": true})
}

func (handler Handler) DeleteInterest(w http.ResponseWriter, r *http.Request) {
	offerID := getOfferID(r)

	err := handler.market.RemoveIntrest(offerID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	render.JSON(w, r, map[string]interface{}{"success": true})
}

func getOfferID(r *http.Request) int64 {
	offerID, _ := strconv.ParseInt(chi.URLParam(r, "offerID"), 10, 64)
	return offerID
}
