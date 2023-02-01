package coffeeco

import "github.com/Rhymond/go-money"

// Product is a value object that models a CoffeeCo product.
type Product struct {
	ItemName  string
	BasePrice money.Money
}
