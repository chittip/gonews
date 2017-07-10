package main

import (
	"net/http"

	"github.com/chittip/gonews/pkg/app"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}
