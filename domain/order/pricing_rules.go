package order

import "github.com/Rhymond/go-money"

type PricingRules interface {
	price(orderItems OrderItems) *money.Money
}