package pricing

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
)

type SimplePricingRules struct {
	rules []order.DiscountRule
}

func New(rules []order.DiscountRule) *SimplePricingRules {
	return &SimplePricingRules{rules}
}

func (pricingRules SimplePricingRules) price(orderItems order.OrderItems) (totalDiscountedPrice *money.Money, err error) {
	discounts := []order.Discount{}
	for _, pricingRule := range pricingRules.rules {
		discount := pricingRule.Discount(orderItems)
		if discount != nil {
			discounts = append(discounts, *discount)
		}
	}
	totalPriceOfItems := orderItems.TotalPriceOfItems()
	totalDiscountedPrice, err = applyDiscounts(discounts, totalPriceOfItems)
	return
}

func applyDiscounts(discounts []order.Discount, totalPrice *money.Money) (discountedPrice *money.Money, err error) {
	discountedPrice = totalPrice
	for _, discount := range discounts {
		discountedPrice, err = discountedPrice.Subtract(&discount.Amount)
		if err != nil {
			return
		}
	}
	return
}


