package large_quantity

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
	"testing"
	"github.com/jisfy/go-checkout/domain/rules/discounts/fixtures"
)

func TestLargeQuantityDiscountRule_EffectiveDiscountPercentage_WhenThresholdOvercome(t *testing.T) {
	const largeQuantityDiscountPercentage = 30
	const expectedEffectiveDiscountPercentage = 30
	discountRule := &LargeQuantityDiscountRule{2, largeQuantityDiscountPercentage, order.Voucher}
	effectiveDiscountPercentage := discountRule.effectiveDiscountPercentage(4)
	if effectiveDiscountPercentage != expectedEffectiveDiscountPercentage {
			t.Errorf("unexpected effective discountRule percentage. Expected %d, found %d",
				expectedEffectiveDiscountPercentage, effectiveDiscountPercentage)
	}
}

func TestLargeQuantityDiscountRule_EffectiveDiscountPercentage_WhenThresholdNotOvercome(t *testing.T) {
	const largeQuantityDiscountPercentage = 30
	const expectedEffectiveDiscountPercentage = DefaultDiscountPercentage
	discountRule := &LargeQuantityDiscountRule{2, largeQuantityDiscountPercentage, order.Voucher}
	effectiveDiscountPercentage := discountRule.effectiveDiscountPercentage(1)
	if effectiveDiscountPercentage != expectedEffectiveDiscountPercentage {
		t.Errorf("unexpected effective discountRule percentage. Expected %d, found %d",
			expectedEffectiveDiscountPercentage, effectiveDiscountPercentage)
	}
}

func TestLargeQuantityDiscountRule_TotalDiscountedPriceOfItems(t *testing.T) {
	const effectiveDiscountPercentage = 20
	expectedTotalDiscountedPrice := money.New(80, fixtures.EuroCurrencyCode)
	totalPriceOfItems := *money.New(400, fixtures.EuroCurrencyCode)
	discountedPrice := totalDiscountedPriceOfItems(effectiveDiscountPercentage, totalPriceOfItems)
	if equalDiscountedPrice, _ := discountedPrice.Equals(expectedTotalDiscountedPrice); !equalDiscountedPrice {
		t.Errorf("unexpected Discount price. Expected %d, found %d",
			expectedTotalDiscountedPrice.Amount(), discountedPrice.Amount())
	}
}

func TestLargeQuantityDiscountRule_Discount(t *testing.T) {
	const largeQuantityDiscountPercentage = 30
	expectedTotalDiscount := money.New(180, fixtures.EuroCurrencyCode)
	orderItems := fixtures.NewOrderItems(3, 1)
	discountRule := &LargeQuantityDiscountRule{2, largeQuantityDiscountPercentage, order.TShirt}
	discount := discountRule.Discount(orderItems)
	if discount == nil {
		t.Fatalf("unexpected Discount. Expected non nil Discount")
	}
	if areEqualAmounts, _ := discount.Amount.Equals(expectedTotalDiscount); !areEqualAmounts {
		t.Errorf("unexpected Discount amount. Expected %d, found %d",
			expectedTotalDiscount.Amount(), discount.Amount.Amount())
	}
	if discount.Id != LargeQuantityDiscountId {
		t.Errorf("unexpected Discount id. Expected %s, found %s",
			LargeQuantityDiscountId, discount.Id)
	}
}

func TestLargeQuantityDiscountRule_Discount_None(t *testing.T) {
	const largeQuantityDiscountPercentage = 30
	orderItems := fixtures.NewOrderItems(1, 1)
	discountRule := &LargeQuantityDiscountRule{2, largeQuantityDiscountPercentage, order.TShirt}
	discount := discountRule.Discount(orderItems)
	if discount != nil {
		t.Fatalf("unexpected Discount. Expected nil Discount")
	}
}