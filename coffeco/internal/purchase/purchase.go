package purchase

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"
	"context"
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// Purchase is an entity that represents a unique transaction between a CoffeeLover and a Store.
type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	total              money.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}

// validateAndEnrich is a service that represents a purchase. It
// calculates the price of each Product in the purchase, then assigns
// it a unique identifier and timestamps the purchase.
func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("purchase must consist of at least one product")
	}

	p.total = *money.New(0, "USD")

	for _, v := range p.ProductsToPurchase {
		newTotal, _ := p.total.Add(&v.BasePrice)
		p.total = *newTotal
	}

	if p.total.IsZero() {
		return errors.New("please validate--purchase should never be zero")
	}

	p.id = uuid.New()
	p.timeOfPurchase = time.Now()

	return nil
}

// CardChargeService represents a transaction when the means is a charge card.
type CardChargeService interface {
	ChargeCard(ctx context.Context, amount money.Money, cardToken string) error
}

// Service charges a card and stores the interaction in the Repository
type Service struct {
	cardService  CardChargeService
	purchaseRepo Repository
}

// CompletePurchase processes a purchase transaction and saves the transaction
// details in the Repository.
func (s *Service) CompletePurchase(ctx context.Context,
	purchase *Purchase) error {
	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}

	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		if err := s.cardService.ChargeCard(ctx, purchase.total, *purchase.CardToken); err != nil {
			return errors.New("card charge failed, cancelling purchase")
		}
	case payment.MEANS_CASH:
		// TODO
	default:
		return errors.New("unknown payment type")
	}

	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return errors.New("failed to store purchase")
	}
	return nil
}
