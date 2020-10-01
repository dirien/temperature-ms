package handler

import (
	"encoding/json"
	"net/http"

	"it.schwarz/country/app/model"
	"it.schwarz/country/config"

	guuid "github.com/google/uuid"
)

var countries []*model.Country

func initCountryList() {
	countries = countries[:0]
	countries = append(countries, model.NewCountry(guuid.New().String(), "Germany", "DE"))
	countries = append(countries, model.NewCountry(guuid.New().String(), "Bulgaria", "BG"))
	countries = append(countries, model.NewCountry(guuid.New().String(), "France", "FR"))

}

// GetcityMember get city members
func GetCountries(config *config.Config, w http.ResponseWriter, r *http.Request) {
	initCountryList()
	ResponseWriter(w, http.StatusOK, "", countries)

}

// ResponseWriter will write result in http.ResponseWriter
func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(data)
	return err
}
