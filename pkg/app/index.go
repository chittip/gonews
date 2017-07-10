package app

import (
	"net/http"

	"github.com/chittip/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
