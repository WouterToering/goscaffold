package api

import (
	"encoding/json"
	"net/http"
)

func (a *Server) GetPotato() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		potato := a.Services.PotatoService.GetPotato(name)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(potato)
	}
}
