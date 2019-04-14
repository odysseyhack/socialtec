package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
)

var (
	routes = flag.Bool("routes", false, "Generate api documentation")
	key    = flag.String("addr", "536d2fffff9af2dcb66e75782ccf75450246703130b8ab775f1f5893a6cef26a", "Node address")
	port   = flag.Int("port", 3333, "http port")
)

func main() {
	flag.Parse()
	r := chi.NewRouter()
	initDependencies(&Config{NodeKey: *key})
	addMiddleWares(r)
	addRouters(r)
	if *routes {
		fmt.Print(docgen.JSONRoutesDoc(r))
		return
	}
	http.Handle("/static", http.FileServer(http.Dir("static")))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), r); err != nil {
		fmt.Println("failed listening:", err)
	}
}
