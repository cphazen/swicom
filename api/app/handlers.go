package app

import (
	"encoding/json"
	"example/swift-comply/database/mapping"
	"example/swift-comply/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		res, err := a.DB.GetCats()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var resp = make([]models.CatJson, len(res))
		for idx, cat := range res {
			resp[idx] = mapping.CatJson(cat)
		}
		jsonResponse(w, r, http.StatusOK, resp)
	}
}

func (a *App) GetCatHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := mux.Vars(r)["id"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res, err := a.DB.GetCatById(idInt)
		fmt.Println("res", res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if res == (models.Cat{}) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		jsonResponse(w, r, http.StatusOK, mapping.CatJson(res))
	}
}
