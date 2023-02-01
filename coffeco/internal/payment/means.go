package payment

// Means represents how the CoffeeLover purchases a product.
type Means string

const (
	MEANS_CARD      = "card"
	MEANS_CASH      = "cash"
	MEANS_COFFEEBUX = "coffeebux"
)

// CardDetails represents a credit card and its number.
type CardDetails struct {
	cardToken string
}
