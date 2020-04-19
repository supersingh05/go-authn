package responses

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}
