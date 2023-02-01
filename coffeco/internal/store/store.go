package store

import (
	coffeeco "coffeeco/internal"

	"github.com/google/uuid"
)

// Store is an entity that represents a specific CoffeeCo store.
type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeeco.Product
}
