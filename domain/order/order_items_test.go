package order

import (
	"github.com/Rhymond/go-money"
	"testing"
)

const (
	euroCurrencySymbol = "EUR"
	cabifyTshirtBarcode = 1111
	cabifyMugBarcode = 1112
	cheapTshirtBarcode = 1113
	cabifyMugPriceCentsEuro = 200
	cabifyTshirtPriceCentsEuro = 200
	cheapTshirtPriceCentsEuro = 100
)

var (
	cabifyTShirt Item = Item{cabifyTshirtBarcode, "Cabify TShirt", *money.New(cabifyTshirtPriceCentsEuro, euroCurrencySymbol), TShirt}
	cabifyMug Item = Item{cabifyMugBarcode, "Cabify Mug", *money.New(cabifyMugPriceCentsEuro, euroCurrencySymbol), Mug}
	cheapTShirt Item = Item{cheapTshirtBarcode, "Cheap Tshirt", *money.New(cheapTshirtPriceCentsEuro, euroCurrencySymbol), TShirt}
)

func TestOrderItems_NumberOfItemsWithCode(t *testing.T) {
	var numberOfTShirts uint = 2
	var numberOfMugs uint = 1

	testDataTable := []struct {
		itemCode     ItemCode
		itemCodeName string
		expected     uint
	}{{TShirt, "TShirt", 2}, {Voucher, "Voucher", 0}}

	orderItems := newOrderItems(numberOfTShirts, numberOfMugs)

	for _, testData := range testDataTable {
		itemsWithCode := orderItems.NumberOfItemsWithCode(testData.itemCode)
		if itemsWithCode != testData.expected {
			t.Errorf("Unexpected number of Items with Code %s, expected %d, found %d",
				testData.itemCodeName, testData.expected, itemsWithCode)
		}
	}
}

func newOrderItems(numberOfTShirts, numberOfMugs uint) OrderItems {
	tshirtOrderItem := OrderItem{cabifyTShirt, numberOfTShirts }
	mugOrderItem := OrderItem{cabifyMug, numberOfMugs }
	return OrderItems{[]OrderItem{tshirtOrderItem, mugOrderItem}}
}

func TestOrderItems_NumberOfItemsWithCode_Empty(t *testing.T) {
	orderItems := OrderItems{}

	itemsWithCode := orderItems.NumberOfItemsWithCode(Voucher)
	if itemsWithCode != 0 {
		t.Errorf("Unexpected number of Items with Code expected %d, found %d",0, itemsWithCode)
	}
}

func TestOrderItems_TotalPriceOfItems(t *testing.T) {
	const numberOfTShirts = 2
	const numberOfMugs = 1
	testDataTable := []struct {
		itemCode ItemCode
		expected money.Money
	}{
		{TShirt,*money.New(400, euroCurrencySymbol)},
		{Voucher,*money.New(0, euroCurrencySymbol)},
	}
	orderItems := newOrderItems(numberOfTShirts, numberOfMugs)

	for _, testData := range testDataTable {
		totalPriceOfItems := orderItems.TotalPriceOfItemsByCode(testData.itemCode)
		if areEqual, _ := testData.expected.Equals(totalPriceOfItems); !areEqual {
			t.Errorf("Unexpected total Price of Items, expected %d found %d",
				totalPriceOfItems.Amount(), testData.expected.Amount())
		}
	}
}

func TestOrderItems_TotalPriceOfItems_Empty(t *testing.T) {
	orderItems := OrderItems{}

	totalPriceOfItems := orderItems.TotalPriceOfItemsByCode(Voucher)
	if !totalPriceOfItems.IsZero() {
		t.Errorf("Unexpected total Price of Items, expected %d found %d",0, totalPriceOfItems.Amount())
	}
}

func TestOrderItems_MinPriceOfItems(t *testing.T) {
	const numberOfTShirts = 2

	tshirtOrderItem := OrderItem{cabifyTShirt, numberOfTShirts }
	cheaperTShirtOrderItem := OrderItem{cheapTShirt, numberOfTShirts }
	orderItems := OrderItems{[]OrderItem{tshirtOrderItem, cheaperTShirtOrderItem}}

	minPriceOfItems, _ := orderItems.MinPriceOfItems(TShirt)
	if minPriceOfItems.Amount() != cheapTshirtPriceCentsEuro {
		t.Errorf("Unexpected minimum Price of Items, expected %d found %d", cheapTshirtPriceCentsEuro, minPriceOfItems.Amount())
	}
}

func TestOrderItems_MinPriceOfItems_Error(t *testing.T) {
	const numberOfTShirts = 2

	tshirtOrderItem := OrderItem{cabifyTShirt, numberOfTShirts }
	cheaperTShirtOrderItem := OrderItem{cheapTShirt, numberOfTShirts }
	orderItems := OrderItems{[]OrderItem{tshirtOrderItem, cheaperTShirtOrderItem}}

	_, error := orderItems.MinPriceOfItems(Voucher)
	if error == nil {
		t.Errorf("Unexpected minimum Price of Items, expected error ")
	}
}