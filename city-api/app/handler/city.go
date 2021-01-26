package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	guuid "github.com/google/uuid"
	"it.schwarz/city/app/model"
	"it.schwarz/city/config"
)

var cities []*model.City

func initCityList() {
	cities = cities[:0]
	cities = append(cities, model.NewCity(guuid.New().String(), "Berlin", 10, "DE"))
	cities = append(cities, model.NewCity(guuid.New().String(), "Stuttgart", 15, "DE"))

	cities = append(cities, model.NewCity(guuid.New().String(), "Sofia", 11, "BG"))
	cities = append(cities, model.NewCity(guuid.New().String(), "Varna", 8, "BG"))

	cities = append(cities, model.NewCity(guuid.New().String(), "Paris", 21, "FR"))
	cities = append(cities, model.NewCity(guuid.New().String(), "Lyon", 28, "FR"))
}

func filter(vs []*model.City, f func(string) bool) []*model.City {
	vsf := make([]*model.City, 0)
	for _, v := range vs {
		if f(v.Country) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// GetcityMember get city members
func GetCities(config *config.Config, w http.ResponseWriter, r *http.Request) {
	if len(cities) == 0 {
		initCityList()
	}
	query := r.URL.Query().Get("country")
	if len(query) > 0 {
		ccities := filter(cities, func(v string) bool {
			return strings.Contains(v, query)
		})
		ResponseWriter(w, http.StatusOK, ccities)
	} else {
		ResponseWriter(w, http.StatusOK, cities)
	}
}

func AddNewCity(config *config.Config, w http.ResponseWriter, r *http.Request) {
	// Declare a new City struct.
	var newCity *model.City

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&newCity)
	if err != nil {
		ResponseWriter(w, http.StatusBadRequest, "BAD REQUEST")
		return
	}

	newCity.ID = guuid.New().String()
	// Do something with the City struct...
	cities = append(cities, newCity)
	ResponseWriter(w, http.StatusCreated, newCity)
}

// ResponseWriter will write result in http.ResponseWriter
func ResponseWriter(res http.ResponseWriter, statusCode int, data interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(data)
	return err
}
