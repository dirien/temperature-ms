package model

type (

	// Member for user in landmark
	Landmark struct {
		ID      string `json:"id"`
		Name    string `json:"name" `
		Country string `json:"country"`
	}
)

// NewADOProject creates a new NewADOProject
func NewLandmark(id string, name string, country string) *Landmark {
	return &Landmark{
		ID:      id,
		Name:    name,
		Country: country,
	}
}
