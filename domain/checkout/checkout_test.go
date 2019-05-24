package checkout

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
	"github.com/jisfy/go-checkout/domain/rules/discounts/fixtures"
	"github.com/jisfy/go-checkout/domain/rules/discounts/large_quantity"
	"github.com/jisfy/go-checkout/domain/rules/discounts/take_n_pay_m"
	"github.com/jisfy/go-checkout/domain/rules/pricing"
	"testing"
)

func newCheckout() *Checkout {
	twoVouchersForOneRule, _ := take_n_pay_m.New(2, 1, order.Voucher)
	fivePercentOffOverThreeShirts :=
		&large_quantity.LargeQuantityDiscountRule{3, 5, order.TShirt}
	discountRules := []order.DiscountRule{twoVouchersForOneRule, fivePercentOffOverThreeShirts}
	return New(pricing.New(discountRules))
}

func TestCheckout_WhenVoucherShirtMugAreScannedThenTotalReturnsCorrectResult(t *testing.T) {
	checkout := newCheckout()
	checkout.scan(fixtures.NewVoucher())
	checkout.scan(fixtures.NewTShirt())
	checkout.scan(fixtures.NewMug())
	total, err := checkout.total()
	if err != nil {
		t.Fatalf("unexpected error performing checkout total %v \n", err)
	}
	expectedCheckoutTotal := money.New(3250, fixtures.EuroCurrencyCode)
	if areEqual, _ := total.Equals(expectedCheckoutTotal); !areEqual {
		t.Fatalf("unexpected total checkout amount. Expected %d found %d \n",
			expectedCheckoutTotal.Amount(), total.Amount())
	}
}

func TestCheckout_WhenVoucherShirtVoucherAreScannedThenTotalReturnsCorrectResult(t *testing.T) {
	checkout := newCheckout()
	checkout.scan(fixtures.NewVoucher())
	checkout.scan(fixtures.NewTShirt())
	checkout.scan(fixtures.NewVoucher())
	total, err := checkout.total()
	if err != nil {
		t.Fatalf("unexpected error performing checkout total %v \n", err)
	}
	expectedCheckoutTotal := money.New(2500, fixtures.EuroCurrencyCode)
	if areEqual, _ := total.Equals(expectedCheckoutTotal); !areEqual {
		t.Fatalf("unexpected total checkout amount. Expected %d found %d \n",
			expectedCheckoutTotal.Amount(), total.Amount())
	}
}

func TestCheckout_WhenThreeShirtsVoucherShirtAreScannedThenTotalReturnsCorrectResult(t *testing.T) {
	checkout := newCheckout()
	checkout.scan(fixtures.NewTShirt())
	checkout.scan(fixtures.NewTShirt())
	checkout.scan(fixtures.NewTShirt())
	checkout.scan(fixtures.NewVoucher())
	checkout.scan(fixtures.NewTShirt())
	total, err := checkout.total()
	if err != nil {
		t.Fatalf("unexpected error performing checkout total %v \n", err)
	}
	expectedCheckoutTotal := money.New(8100, fixtures.EuroCurrencyCode)
	if areEqual, _ := total.Equals(expectedCheckoutTotal); !areEqual {
		t.Fatalf("unexpected total checkout amount. Expected %d found %d \n",
			expectedCheckoutTotal.Amount(), total.Amount())
	}
}

func TestCheckout_WhenVoucherShirtTwoVouchersMugAndTwoShirtsAreScannedThenTotalReturnsCorrectResult(t *testing.T) {
	checkout := newCheckout()
	checkout.scan(fixtures.NewVoucher())
	checkout.scan(fixtures.NewTShirt())
	checkout.scan(fixtures.NewVoucher())
	checkout.scan(fixtures.NewVoucher())
	checkout.scan(fixtures.NewMug())
	checkout.scan(fixtures.NewTShirt())
	checkout.scan(fixtures.NewTShirt())
	total, err := checkout.total()
	if err != nil {
		t.Fatalf("unexpected error performing checkout total %v \n", err)
	}
	expectedCheckoutTotal := money.New(7450, fixtures.EuroCurrencyCode)
	if areEqual, _ := total.Equals(expectedCheckoutTotal); !areEqual {
		t.Fatalf("unexpected total checkout amount. Expected %d found %d \n",
			expectedCheckoutTotal.Amount(), total.Amount())
	}
}
