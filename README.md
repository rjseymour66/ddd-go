# Domain Driven Design with Go

1. Define the ubiquitous language.
2. Identify the domains.
3. Required functionality for a minimum viable product (MVP)? User stories(?)

## Concepts

### Value objects

Value objects are not unique. When deciding whether to model something as a value object, ask the following questions:
- Is this object immutable?
- Does it measure, quantify, or describe a domain concept?
- Can it be compared to other objects of the same type by just its values?

If you are still unsure, you can treat something as a value object and then upgrade it to an entity later.

```go
// Product is a value object that models a CoffeeCo product.
type Product struct {
	ItemName  string
	BasePrice money.Money
}
```

### Services

After you complete the domain models, you can start building the services. The following describe a service:
- It performs a significant piece of business logic within out domain.
- It calculates some values.
- It interacts with the repository layer.

Try to push down as much logic as possible into your domain objects (i.e. create service functions within the domain `.go` files.)


## Project structure

- `/internal` directory is special because other projects cannot import anything  in this directory. This is a good place for domain code.
  If an entity needs to be accessible to all domain code, place its file in the `/internal` directory. Its package name is the project package.
  Each subdirectory in `/internal` is a domain.