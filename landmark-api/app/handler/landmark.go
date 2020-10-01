package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	guuid "github.com/google/uuid"
	"it.schwarz/landmark/app/model"
	"it.schwarz/landmark/config"
)

var landmarks []*model.Landmark

func initLandmarkList() {
	landmarks = landmarks[:0]
	landmarks = append(landmarks, model.NewLandmark(guuid.New().String(), "Berlin Wall Monument", "DE"))
	landmarks = append(landmarks, model.NewLandmark(guuid.New().String(), "Reichstag", "DE"))

	landmarks = append(landmarks, model.NewLandmark(guuid.New().String(), "Seven Rila Lakes", "BG"))
	landmarks = append(landmarks, model.NewLandmark(guuid.New().String(), "The Cave Yagodina", "BG"))

	landmarks = append(landmarks, model.NewLandmark(guuid.New().String(), "Eiffel Tower", "FR"))
	landmarks = append(landmarks, model.NewLandmark(guuid.New().String(), "Lyon", "FR"))
}

func filter(vs []*model.Landmark, f func(string) bool) []*model.Landmark {
	vsf := make([]*model.Landmark, 0)
	for _, v := range vs {
		if f(v.Country) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// GetLandmarkMember get landmark members
func GetLandmarks(config *config.Config, w http.ResponseWriter, r *http.Request) {
	if len(landmarks) == 0 {
		initLandmarkList()
	}
	query := r.URL.Query().Get("country")
	if len(query) > 0 {
		llandmarks := filter(landmarks, func(v string) bool {
			return strings.Contains(v, query)
		})
		ResponseWriter(w, http.StatusOK, llandmarks)
	} else {
		ResponseWriter(w, http.StatusOK, landmarks)
	}
}

func AddNewLandmark(config *config.Config, w http.ResponseWriter, r *http.Request) {
	// Declare a new Landmark struct.
	var newLandmark model.Landmark

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&newLandmark)
	if err != nil {
		ResponseWriter(w, http.StatusBadRequest, "BAD REQUEST")
		return
	}

	newLandmark.ID = guuid.New().String()

	// Do something with the Landmark struct...
	landmarks = append(landmarks, &newLandmark)
	ResponseWriter(w, http.StatusCreated, &newLandmark)

}

// ResponseWriter will write result in http.ResponseWriter
func ResponseWriter(res http.ResponseWriter, statusCode int, data interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	err := json.NewEncoder(res).Encode(data)
	return err
}
