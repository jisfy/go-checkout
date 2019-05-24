package order

import "github.com/Rhymond/go-money"

type PricingRules interface {
	Price(orderItems OrderItems) (*money.Money, error)
}
