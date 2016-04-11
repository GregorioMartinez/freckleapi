package freckleapi

type Tag struct {
	Billable      bool    `json:"billable"`
	FormattedName string  `json:"formatted_name"`
	ID            float64 `json:"id"`
	Name          string  `json:"name"`
	URL           string  `json:"url"`
}
