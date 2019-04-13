package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"

	"github.com/odysseyhack/socialtec/node/cmd/handlers"
)

var routes = flag.Bool("routes", false, "Generate api documentation")

func addRouters(r *chi.Mux) {
	handler := handlers.NewHandler(nil)

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
	r.Put("/interest/{offerID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("intrest added"))
	})

	// DELETE /interest/:product_id  remove intrest in an item
	r.Delete("/interest/{offerID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("intrest delete"))
	})
}

func addMiddleWares(r *chi.Mux) {
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

func initDependencies(conf *Config) error {
	return nil
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
	http.ListenAndServe(":3333", r)
}
