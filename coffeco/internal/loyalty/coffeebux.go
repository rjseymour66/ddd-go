package loyalty

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/store"

	"github.com/google/uuid"
)

// CoffeeBux is an entity that represents a specific account in the loyalty program.
type CoffeeBux struct {
	ID                                   uuid.UUID
	store                                store.Store
	coffeeLover                          coffeeco.CoffeeLover
	FreeDrinksAvailable                  int
	RemaininDrivePurchasesUntilFreeDrink int
}
