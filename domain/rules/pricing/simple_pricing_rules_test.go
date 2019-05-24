package pricing

import (
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
	"github.com/jisfy/go-checkout/domain/rules/discounts/fixtures"
	"github.com/jisfy/go-checkout/domain/rules/discounts/large_quantity"
	"github.com/jisfy/go-checkout/domain/rules/discounts/take_n_pay_m"
	"testing"
)

func TestSimplePricingRules_ApplyDiscounts(t *testing.T) {
	takeThreeTShirtsPayTwoDiscount :=
		order.Discount{take_n_pay_m.TakeMPayNDiscountId, *money.New(200, fixtures.EuroCurrencyCode)}
	twentyPercentOffMoreThanFourMugs :=
		order.Discount{large_quantity.LargeQuantityDiscountId, *money.New(160, fixtures.EuroCurrencyCode)}
	discounts := []order.Discount{takeThreeTShirtsPayTwoDiscount, twentyPercentOffMoreThanFourMugs}
	totalPriceOfThreeShirtsAndFourMugs := money.New((fixtures.CabifyTshirtPriceCentsEuro*3)+
		(fixtures.CabifyMugPriceCentsEuro*4), fixtures.EuroCurrencyCode)
	expectedTotalDiscountedPrice := money.New(
		(fixtures.CabifyTshirtPriceCentsEuro*2)+(fixtures.CabifyMugPriceCentsEuro*3+40), fixtures.EuroCurrencyCode)

	totalDiscountedPrice, err := applyDiscounts(discounts, totalPriceOfThreeShirtsAndFourMugs)

	if err != nil {
		t.Fatalf("unexpected error applying discounts %v \n", err)
	}
	if areEqual, _ := totalDiscountedPrice.Equals(expectedTotalDiscountedPrice); !areEqual {
		t.Fatalf("unexpected value of total discounted Price. Expected %d found %d",
			expectedTotalDiscountedPrice.Amount(), totalDiscountedPrice.Amount())
	}
}

func TestSimplePricingRules_Price(t *testing.T) {
	testData := []struct {
		itemsToPay            uint
		itemstoTake           uint
		itemsThreshold        uint
		percentageOff         uint
		numberOfOfTShirts     uint
		numberOfMugs          uint
		expectedDiscountPrice *money.Money
	}{
		{2, 3, 4, 20, 4, 1, money.New(
			(fixtures.CabifyTshirtPriceCentsEuro*3)+fixtures.CabifyMugPriceCentsEuro, fixtures.EuroCurrencyCode)},
		{2, 3, 4, 20, 4, 4, money.New(
			(fixtures.CabifyTshirtPriceCentsEuro*3)+(fixtures.CabifyMugPriceCentsEuro*3+40), fixtures.EuroCurrencyCode)},
		{2, 3, 4, 20, 2, 1, money.New(
			(fixtures.CabifyTshirtPriceCentsEuro*2)+fixtures.CabifyMugPriceCentsEuro, fixtures.EuroCurrencyCode)},
		{2, 3, 4, 20, 2, 0, money.New(fixtures.CabifyTshirtPriceCentsEuro*2, fixtures.EuroCurrencyCode)},
	}

	for _, test := range testData {
		takeThreeTShirtsPayTwoDiscountRule, _ := take_n_pay_m.New(test.itemstoTake, test.itemsToPay, order.TShirt)
		twentyPercentOffFourMugsOrMoreDiscountRule :=
			&large_quantity.LargeQuantityDiscountRule{test.itemsThreshold, test.percentageOff, order.Mug}

		orderItems := fixtures.NewOrderItems(test.numberOfOfTShirts, test.numberOfMugs)

		simplePricingRules :=
			New([]order.DiscountRule{takeThreeTShirtsPayTwoDiscountRule, twentyPercentOffFourMugsOrMoreDiscountRule})

		totalDiscountedPrice, err := simplePricingRules.Price(orderItems)

		if err != nil {
			t.Fatalf("unexpected error applying discounts %v \n", err)
		}
		if areEqual, _ := totalDiscountedPrice.Equals(test.expectedDiscountPrice); !areEqual {
			t.Fatalf("unexpected value of total discounted Price. Expected %d found %d",
				test.expectedDiscountPrice.Amount(), totalDiscountedPrice.Amount())
		}
	}
}
