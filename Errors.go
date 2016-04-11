package freckleapi

//@TODO Check if this is even used, yo
type Errors struct {
	Message string `json:"message"`
}

func (e Errors) Error() string {
	return e.Message
}

type NoMorePages struct {
	Message string `json:"message"`
}

func (e NoMorePages) Error() string {
	return e.Message
}
