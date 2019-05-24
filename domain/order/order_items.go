package order

import (
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
)

const defaultCurrencyCode = "EUR"

type OrderItems struct {
	Items []OrderItem
}

func (orderItems OrderItems) TotalPriceOfItems() *money.Money {
	var totalPriceOfItems *money.Money = money.New(0, defaultCurrencyCode)
	for _, orderItem := range orderItems.Items {
		totalPriceOfItems, _ = totalPriceOfItems.Add(orderItem.totalPrice())
	}
	return totalPriceOfItems
}

func (orderItems OrderItems) TotalPriceOfItemsByCode(itemCode ItemCode) *money.Money {
	var totalPriceOfItems *money.Money = money.New(0, defaultCurrencyCode) // should replace default currency Code
	for _, orderItem := range orderItems.Items {
		if orderItem.Item.Code == itemCode {
			totalPriceOfItems, _ = totalPriceOfItems.Add(orderItem.totalPrice())
		}
	}
	return totalPriceOfItems
}

func (orderItems OrderItems) MinPriceOfItems(itemCode ItemCode) (minItemPrice *money.Money, err error) {
	minItemPrice = nil
	err = nil
	for _, orderItem := range orderItems.Items {
		if orderItem.Item.Code == itemCode {
			if minItemPrice == nil {
				minItemPrice = money.New(orderItem.Item.Price.Amount(), orderItem.Item.Price.Currency().Code)
			}
			if orderItemPricesIsSmaller, _ := orderItem.Item.Price.LessThan(minItemPrice); orderItemPricesIsSmaller {
				minItemPrice = money.New(orderItem.Item.Price.Amount(), orderItem.Item.Price.Currency().Code)
			}
		}
	}
	if minItemPrice == nil {
		err = errors.New(fmt.Sprintf("Could not find the minimum Item Price of type %d", itemCode))
	}
	return
}

func (orderItems OrderItems) NumberOfItemsWithCode(itemCode ItemCode) uint {
	var itemsWithCode uint = 0
	for _, orderItem := range orderItems.Items {
		if orderItem.Item.Code == itemCode {
			itemsWithCode += orderItem.Quantity
		}
	}
	return itemsWithCode
}
