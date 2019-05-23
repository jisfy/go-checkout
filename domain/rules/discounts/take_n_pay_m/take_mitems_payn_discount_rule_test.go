package take_n_pay_m

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
	"github.com/jisfy/go-checkout/domain/rules/discounts/fixtures"
	"testing"
)

func TestTakeMItemsPayNWithCodeDiscountRule_New_InvalidItemsToPay(t *testing.T) {
	const numberOfOItemsToTake = 3
	const numberOfItemsToPay = 0
	_, err := New(numberOfOItemsToTake, numberOfItemsToPay, order.TShirt)
	if err == nil {
		t.Fatalf("expected non nil error, got a nil one")
	}
}

func TestTakeMItemsPayNWithCodeDiscountRule_New_ValidItemsToPay(t *testing.T) {
	const numberOfOItemsToTake = 3
	const numberOfItemsToPay = 1
	_, err := New(numberOfOItemsToTake, numberOfItemsToPay, order.TShirt)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestTakeMItemsPayNWithCodeDiscountRule_TotalNumberOfFreeItems(t *testing.T) {
	testData := []struct {
		itemsToPay uint
		itemstoTake uint
		itemsWithCode uint
		expectedFreeItems uint
	}{{2, 3, 4, 1}, {2, 3, 6, 2}, {2, 3, 2, 0}}

	for _, test := range testData {
		discountRule, _ := New(test.itemstoTake, test.itemsToPay, order.TShirt)
		totalNumberOfFreeItems := discountRule.totalNumberOfFreeItems(test.itemsWithCode)
		if totalNumberOfFreeItems != test.expectedFreeItems {
			t.Errorf("unexpected number of free items. Expected %d, found %d",
				test.expectedFreeItems, totalNumberOfFreeItems)
		}
	}
}

func TestTakeMItemsPayNWithCodeDiscountRule_TotalPriceOfFreeItems(t *testing.T) {
	testData := []struct {
		itemsToPay uint
		itemstoTake uint
		numberOfFreeItems uint
		numberOfItemsOfType uint
		expectedFreeItemsPrice *money.Money
		expectsError bool
	}{
		{2, 3, 1, 4, money.New(200, fixtures.EuroCurrencyCode), false},
		{2, 3, 2, 6, money.New(400, fixtures.EuroCurrencyCode), false},
		{2, 3, 0, 0, nil, true},
	}

	for _, test := range testData {
		orderItems := fixtures.NewOrderItems(test.numberOfItemsOfType, 1)
		discountRule, err := New(test.itemstoTake, test.itemsToPay, order.TShirt)
		totalPriceOfFreeItems, err := discountRule.totalPriceOfFreeItems(orderItems, test.numberOfFreeItems)
		if err != nil && !test.expectsError {
			t.Fatalf("unexpected error found when calculating the total price of free items %v", err)
		}
		if err == nil {
			if test.expectsError {
				t.Fatalf("expecting error that wasn't found %v", totalPriceOfFreeItems.Amount())
			}
			if areEqual, _ := totalPriceOfFreeItems.Equals(test.expectedFreeItemsPrice); !areEqual {
				t.Errorf("unexpected total price of free items. Expected %d, found %d",
					test.expectedFreeItemsPrice.Amount(), totalPriceOfFreeItems.Amount())
			}
		}
	}
}

func TestTakeMItemsPayNWithCodeDiscountRule_Discount(t *testing.T) {
	const numberOfOItemsToTake = 3
	const numberOfItemsToPay = 2
	expectedTotalDiscount := money.New(200, fixtures.EuroCurrencyCode)
	orderItems := fixtures.NewOrderItems(4, 1)

	discountRule, _ := New(numberOfOItemsToTake, numberOfItemsToPay, order.TShirt)
	discount := discountRule.Discount(orderItems)

	if discount == nil {
		t.Fatalf("unexpected Discount. Expected non nil Discount")
	}
	if areEqualAmounts, _ := discount.Amount.Equals(expectedTotalDiscount); !areEqualAmounts {
		t.Errorf("unexpected Discount amount. Expected %d, found %d",
			expectedTotalDiscount.Amount(), discount.Amount.Amount())
	}
	if discount.Id != TakeMPayNDiscountId {
		t.Errorf("unexpected Discount id. Expected %s, found %s",
			TakeMPayNDiscountId, discount.Id)
	}
}

func TestTakeMItemsPayNWithCodeDiscountRule_Discount_None(t *testing.T) {
	testData := []struct {
		itemsToPay uint
		itemstoTake uint
		numberOfItemsOfType uint
		itemCode order.ItemCode
		expectedDiscount *order.Discount
	}{
		{2, 3, 2, order.TShirt, nil},
		{2, 3, 2, order.Voucher, nil},
	}

	for _, test := range testData {
		orderItems := fixtures.NewOrderItems(test.numberOfItemsOfType, 1)

		discountRule, _ := New(test.itemstoTake, test.itemsToPay, test.itemCode)
		discount := discountRule.Discount(orderItems)

		if discount != test.expectedDiscount {
			t.Fatalf("unexpected Discount. Expected nil Discount")
		}
	}
}
