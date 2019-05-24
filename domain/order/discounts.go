package order

import "github.com/Rhymond/go-money"

type Discount struct {
	Id     string
	Amount money.Money
}

type DiscountRule interface {
	Discount(orderItems OrderItems) *Discount
}
