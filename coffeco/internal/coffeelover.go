package coffeeco

import "github.com/google/uuid"

// CoffeeLover is an entity that represents a specific CoffeeCo customer.
type CoffeeLover struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	EmailAddress string
}
