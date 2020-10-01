package model

type (

	// Member for user in city
	City struct {
		ID          string `json:"id"`
		Name        string `json:"name" `
		Temperature int    `json:"temperature"`
		Country     string `json:"country"`
	}
)

// NewADOProject creates a new NewADOProject
func NewCity(id string, name string, temperature int, country string) *City {
	return &City{
		ID:          id,
		Name:        name,
		Temperature: temperature,
		Country:     country,
	}
}
