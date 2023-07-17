package app

import (
	"encoding/json"
	"example/swift-comply/database/entities"
	"example/swift-comply/models"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (a *App) GetCatsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := a.DB.GetCats()
		if res == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var resp = make([]models.CatJson, len(res))
		for idx, post := range res {
			resp[idx] = entities.CatJson(post)
		}
		jsonResponse(w, r, http.StatusOK, resp)
	}
}
