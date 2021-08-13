package amqp

type OrderToProcessView struct {
	ID    string `json:"id"`
	Price PriceView
}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}
