package freckleapi

type User struct {
	Email           string  `json:"email"`
	FirstName       string  `json:"first_name"`
	ID              float64 `json:"id"`
	LastName        string  `json:"last_name"`
	ProfileImageURL string  `json:"profile_image_url"`
	URL             string  `json:"url"`
}
