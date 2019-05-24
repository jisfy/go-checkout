package order

import "github.com/Rhymond/go-money"

type OrderItem struct {
	Item     Item
	Quantity uint
}

func (orderItem OrderItem) totalPrice() *money.Money {
	return orderItem.Item.Price.Multiply(int64(orderItem.Quantity))
}

func (orderItem *OrderItem) IncrementQuantity(quantity uint) {
	orderItem.Quantity += quantity
}
