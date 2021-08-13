package http

type productView struct {
	ID string `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Price priceView `json:"price"`
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}
