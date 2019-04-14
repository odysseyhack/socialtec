package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/odysseyhack/socialtec/node/cmd/handlers"
	"github.com/odysseyhack/socialtec/node/pkg/market"
	"github.com/odysseyhack/socialtec/node/pkg/store"
)

func addRouters(r *chi.Mux) {
	handler := handlers.NewHandler(market.DefaultMarket, store.Default)

	// /ping is used as the health check route
	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})

	// GET /offers returns all the available offers
	r.Get("/offers", handler.GetOffers)

	// POST /offers registers items as being offered by the user
	r.Post("/offers", handler.NewOffer)

	// DELETE /offers deletes and offer
	r.Delete("/offers/{offerID}", handler.DeleteOffer)

	// PUT /interest registers intrest for an item
	r.Put("/interest/{offerID}", handler.AddInterest)

	// DELETE /interest/:product_id  remove intrest in an item
	r.Delete("/interest/{offerID}", handler.DeleteInterest)

	r.Get("/me/interests", handler.MyInterests)
	r.Get("/me/offers", handler.MyOffers)

	r.Post("/cycle", handler.InitiateCycle)
	r.Post("/dontLikeAnt/{nodeID}", handler.DontLikeAny)
}
