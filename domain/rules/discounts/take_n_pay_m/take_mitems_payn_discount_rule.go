package take_n_pay_m

import (
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/jisfy/go-checkout/domain/order"
)

const (
	TakeMPayNDiscountId = "000002_M_FOR_N_DISCOUNT"
)

type TakeMItemsPayNWithCodeDiscountRule struct {
	NumberOfItemsToTake uint
	NumberOfItemsToPay  uint
	ItemCode            order.ItemCode
}

func New(numberOfItemsToTake uint, numberOfItemsToPay uint, itemCode order.ItemCode) (*TakeMItemsPayNWithCodeDiscountRule, error) {
	if numberOfItemsToPay <= 0 || numberOfItemsToPay > numberOfItemsToTake {
		return nil, errors.New(
			fmt.Sprintf("the number of items to pay should be a value between zero and %d", numberOfItemsToTake))
	}
	return &TakeMItemsPayNWithCodeDiscountRule{numberOfItemsToTake, numberOfItemsToPay, itemCode}, nil
}

func (discountRule TakeMItemsPayNWithCodeDiscountRule) Discount(orderItems order.OrderItems) *order.Discount {
	numberOfFreeItems := discountRule.totalNumberOfFreeItems(orderItems.NumberOfItemsWithCode(discountRule.ItemCode))
	totalDiscountAmount, err := discountRule.totalPriceOfFreeItems(orderItems, numberOfFreeItems)
	if err != nil || (err == nil && totalDiscountAmount.IsZero()) {
		return nil
	}
	return &order.Discount{TakeMPayNDiscountId, *totalDiscountAmount}
}

func (discountRule TakeMItemsPayNWithCodeDiscountRule) totalPriceOfFreeItems(
	orderItems order.OrderItems, numberOfFreeItems uint) (totalPrice *money.Money, err error) {
	minItemPrice, err := orderItems.MinPriceOfItems(discountRule.ItemCode)
	if err == nil {
		totalPrice = minItemPrice.Multiply(int64(numberOfFreeItems))
	}
	return
}

func (discountRule TakeMItemsPayNWithCodeDiscountRule) totalNumberOfFreeItems(numberOfItemsWithCode uint) uint {
	numberOfTimesFreeItemsGranted := numberOfItemsWithCode / discountRule.NumberOfItemsToTake
	remainingItemsToPay := numberOfItemsWithCode % discountRule.NumberOfItemsToTake
	totalNumberOfItemsToPay := (numberOfTimesFreeItemsGranted * discountRule.NumberOfItemsToPay) + remainingItemsToPay
	totalNumberOfFreeItems := numberOfItemsWithCode - totalNumberOfItemsToPay
	return totalNumberOfFreeItems
}
