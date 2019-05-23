package fixtures

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
)

const (
	EuroCurrencyCode           = "EUR"
	CabifyTshirtBarcode        = 1111
	CabifyMugBarcode           = 1112
	CheapTshirtBarcode         = 1113
	CabifyMugPriceCentsEuro    = 200
	CabifyTshirtPriceCentsEuro = 200
	CheapTshirtPriceCentsEuro  = 100
)

var (
	cabifyTShirt order.Item =
		order.Item{CabifyTshirtBarcode, "Cabify TShirt", *money.New(CabifyTshirtPriceCentsEuro, EuroCurrencyCode), order.TShirt}
	cabifyMug order.Item =
		order.Item{CabifyMugBarcode, "Cabify Mug", *money.New(CabifyMugPriceCentsEuro, EuroCurrencyCode), order.Mug}
	cheapTShirt order.Item =
		order.Item{CheapTshirtBarcode, "Cheap Tshirt", *money.New(CheapTshirtPriceCentsEuro, EuroCurrencyCode), order.TShirt}
)

func NewOrderItems(numberOfTShirts, numberOfMugs uint) order.OrderItems {
	orderItems := []order.OrderItem{}
	var tshirtOrderItem order.OrderItem
	if numberOfTShirts > 0 {
		tshirtOrderItem = order.OrderItem{cabifyTShirt, numberOfTShirts}
		orderItems = append(orderItems, tshirtOrderItem)
	}
	var mugOrderItem order.OrderItem
	if numberOfMugs > 0 {
		mugOrderItem = order.OrderItem{cabifyMug, numberOfMugs}
		orderItems = append(orderItems, mugOrderItem)
	}
	return order.OrderItems{orderItems}
}
