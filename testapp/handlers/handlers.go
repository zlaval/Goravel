package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"goravel"
	"net/http"
)

type Handlers struct {
	App *goravel.Goravel
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering: ", err)
	}
}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering: ", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering: ", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	testData := "hello world"

	h.App.Session.Put(r.Context(), "name", testData)

	testValue := h.App.Session.GetString(r.Context(), "name")

	vars := make(jet.VarMap)
	vars.Set("name", testValue)

	err := h.App.Render.JetPage(w, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering: ", err)
	}
}
