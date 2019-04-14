package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/odysseyhack/socialtec/node/pkg/market"
	"github.com/odysseyhack/socialtec/node/pkg/store"
)

type Offer struct {
	market.Offer     `json:"offer"`
	*market.Interest `json:"interest,omitempty"`
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
	handler.addMyOffer(offer)
	render.JSON(w, r, offer)
}

func (handler Handler) GetOffers(w http.ResponseWriter, r *http.Request) {
	offers, err := handler.market.GetAvailableOffers()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	intrestedParties := make(map[int64]market.Interest)
	if err := handler.store.Get("intrestedParties", &intrestedParties); err != nil && err != store.ErrorNotFound {
		log.Printf("error getting interested parties. err:%+v", err)
		return
	}

	var priorityOffers, normalOffers []Offer
	for _, offer := range offers {
		if interest, ok := intrestedParties[offer.Node]; ok {
			priorityOffers = append(priorityOffers, Offer{Offer: offer, Interest: &interest})
		}
		normalOffers = append(normalOffers, Offer{Offer: offer})
	}

	if err := store.Default.Set("intrestedParties", make(map[int64]market.Interest)); err != nil && err != store.ErrorNotFound {
		log.Printf("error setting interested parties. Err:%+v", err)
		return
	}
	render.JSON(w, r, append(priorityOffers, normalOffers...))
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
	var offer market.Offer
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	offerID := getOfferID(r)

	err := handler.market.ShowIntrest(offerID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	offer.ID = offerID
	handler.addMyIntrest(offer)
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
	handler.removeMyIntrest(offerID)
	render.JSON(w, r, map[string]interface{}{"success": true, "offer_id": offerID})
}

func (handler Handler) InitiateCycle(w http.ResponseWriter, r *http.Request) {
	var offer Offer
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	if err := handler.market.SendFormChain(
		offer.Offer.Node,
		offer.Offer.ID,
		offer.Interest.What,
		offer.Interest.Who); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	handler.removeMyOffer(offer.Offer.ID)

	render.JSON(w, r, map[string]interface{}{"success": true})
}

//DontLikeAny serves  /dont_likeany/:nodeID
func (handler Handler) DontLikeAny(w http.ResponseWriter, r *http.Request) {
	nodeID := getInt(r, "nodeID")
	var offers []market.Offer
	if err := handler.store.Get("my_interests", &offers); err != nil && err != store.ErrorNotFound {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}
	for _, offer := range offers {
		go handler.market.ReferInterest(offer.ID, nodeID)
	}
	render.JSON(w, r, map[string]interface{}{"success": true})
}

func getInt(r *http.Request, key string) int64 {
	intVal, _ := strconv.ParseInt(chi.URLParam(r, key), 10, 64)
	return intVal
}

func getOfferID(r *http.Request) int64 {
	return getInt(r, "offerID")
}
