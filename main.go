package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
)

type Discount struct {
	id string
	amount money.Money
}

type DiscountRule interface {
	discount() Discount
}

func main() {
	fmt.Println("Hello world!!!")
}
