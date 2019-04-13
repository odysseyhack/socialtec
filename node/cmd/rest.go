package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
)

var routes = flag.Bool("routes", false, "Generate api documentation")

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
