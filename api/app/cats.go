package app

import (
	"example/swift-comply/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *App) GetCatsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := a.DB.GetCats()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var resp = make([]models.CatJson, len(res))
		for idx, cat := range res {
			resp[idx] = CatJson(cat)
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
		jsonResponse(w, r, http.StatusOK, CatJson(res))
	}
}

func BreedGroupJson(bg models.BreedGroup) models.BreedGroupJson {
	return models.BreedGroupJson{
		Id:   bg.Id,
		Name: bg.Name,
	}
}

func BreedJson(b models.Breed) models.BreedJson {
	result := models.BreedJson{
		Id:   b.Id,
		Name: b.Name,
	}
	if b.BreedGroup != nil {
		breedGroup := BreedGroupJson(*b.BreedGroup)
		result.BreedGroup = &breedGroup
	}
	return result
}

func CatJson(c models.Cat) models.CatJson {
	result := models.CatJson{
		Id:          c.Id,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
	}
	if c.Breed != nil {
		breed := BreedJson(*c.Breed)
		result.Breed = &breed
	}
	return result
}
