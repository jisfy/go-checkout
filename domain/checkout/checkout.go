package checkout

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
)

type Checkout struct {
	pricingRules        order.PricingRules
	orderItemsByBarcode map[int]*order.OrderItem
}

func New(pricingRules order.PricingRules) *Checkout {
	return &Checkout{pricingRules, make(map[int]*order.OrderItem)}
}

func (checkout *Checkout) scan(item order.Item) {
	if _, present := checkout.orderItemsByBarcode[item.Barcode]; !present {
		checkout.orderItemsByBarcode[item.Barcode] = &order.OrderItem{item, 0}
	}
	checkout.orderItemsByBarcode[item.Barcode].IncrementQuantity(1)
}

func (checkout *Checkout) total() (*money.Money, error) {
	orderItems := []order.OrderItem{}
	for _, orderItem := range checkout.orderItemsByBarcode {
		orderItems = append(orderItems, *orderItem)
	}
	return checkout.pricingRules.Price(order.OrderItems{orderItems})
}
