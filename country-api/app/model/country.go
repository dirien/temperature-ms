package model

type (

	// Member for user in city
	Country struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Iso  string `json:"iso"`
	}
)

// NewADOProject creates a new NewADOProject
func NewCountry(id string, name string, iso string) *Country {
	return &Country{
		ID:   id,
		Name: name,
		Iso:  iso,
	}
}
