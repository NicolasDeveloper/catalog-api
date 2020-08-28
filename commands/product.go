package commands

//ProductInput response to adapter
type ProductInput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

//ProductOutput response to adapter
type ProductOutput struct {
	ID string `json:"id"`
}
