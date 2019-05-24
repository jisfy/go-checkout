package large_quantity

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
)

const (
	DefaultDiscountPercentage = 0
	LargeQuantityDiscountId   = "000000_LARGE_QUANTITY"
)

type LargeQuantityDiscountRule struct {
	NumberOfItemsWithCodeThreshold uint
	DiscountPercentage             uint
	ItemCode                       order.ItemCode
}

func totalDiscountedPriceOfItems(effectiveDiscountPercentage uint,
	totalPriceOfItemsWithCode money.Money) money.Money {
	return *totalPriceOfItemsWithCode.Multiply(int64(effectiveDiscountPercentage)).Divide(100)
}

func (discountRule LargeQuantityDiscountRule) Discount(orderItems order.OrderItems) *order.Discount {
	effectiveDiscountPercentage :=
		discountRule.effectiveDiscountPercentage(orderItems.NumberOfItemsWithCode(discountRule.ItemCode))
	if effectiveDiscountPercentage == 0 {
		return nil
	}
	totalPriceOfItemsWithCode := orderItems.TotalPriceOfItemsByCode(discountRule.ItemCode)
	totalDiscountedPrice := totalDiscountedPriceOfItems(effectiveDiscountPercentage, *totalPriceOfItemsWithCode)
	return &order.Discount{LargeQuantityDiscountId, totalDiscountedPrice}
}

func (discountRule LargeQuantityDiscountRule) effectiveDiscountPercentage(numberOfItemsWithCode uint) uint {
	if numberOfItemsWithCode >= discountRule.NumberOfItemsWithCodeThreshold {
		return discountRule.DiscountPercentage
	}
	return DefaultDiscountPercentage
}
