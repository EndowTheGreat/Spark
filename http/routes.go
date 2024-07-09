package http

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/EndowTheGreat/spark/markdown"

	"github.com/go-chi/chi"
)

func (h *Handler) serveFile(w http.ResponseWriter, r *http.Request) {
	file := chi.URLParam(r, "*")
	if !strings.Contains(file, ".") {
		file += ".html"
	}
	http.ServeFile(w, r, fmt.Sprintf("%v/%v", markdown.Output, file))
}

func (h *Handler) ServeHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("%v/index.html", markdown.Output))
}
