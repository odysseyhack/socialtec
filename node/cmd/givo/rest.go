package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate api documentation")

func addRouters(r *chi.Mux) {
	// /ping is used as the health check route
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// GET /offers returns all the available offers
	r.Get("/offers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("offers returned"))
	})

	// POST /offers registers items as being offered by the user
	r.Post("/offers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("offers added"))
	})

	// PUT /interest registers intrest for an item
	r.Put("/interest/:offer_id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("intrest added"))
	})

	// DELETE /intrest/:product_id  remove intrest in an item
	r.Delete("/intrest/:offer_id", func(w http.ResponseWriter, r *http.Request) {
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
}

func initDependencies(conf *Config) error {
	return nil
}

func main() {
	flag.Parse()

	r := chi.NewRouter()
	addMiddleWares(r)
	addRouters(r)

	if *routes {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/odysseyhack/socialtec/api/cmd/givo",
			Intro:       "Givo API",
		}))
		return
	}
	http.ListenAndServe(":3333", r)
}
