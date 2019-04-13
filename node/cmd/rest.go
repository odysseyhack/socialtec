package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"

	"github.com/odysseyhack/socialtec/node/cmd/handlers"
	"github.com/odysseyhack/socialtec/node/pkg/market"
)

var routes = flag.Bool("routes", false, "Generate api documentation")

func addRouters(r *chi.Mux) {
	handler := handlers.NewHandler(market.DefaultMarket)

	// /ping is used as the health check route
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
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
}

func addMiddleWares(r *chi.Mux) {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
}

// Config defines the structure of config
type Config struct {
	BlockChain struct {
		Host string
		Port string
	}
	Host string
	Port int
}

func main() {
	flag.Parse()
	r := chi.NewRouter()
	initDependencies(&Config{})
	addMiddleWares(r)
	addRouters(r)
	if *routes {
		fmt.Print(docgen.JSONRoutesDoc(r))
		return
	}
	if err := http.ListenAndServe(":3333", r); err != nil {
		fmt.Println("failed listening:", err)
	}
}
