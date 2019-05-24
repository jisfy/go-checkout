package order

import "github.com/Rhymond/go-money"

type ItemCode int

const (
	Voucher = 0
	TShirt  = 1
	Mug     = 2
)

type Item struct {
	Barcode int
	Name    string
	Price   money.Money
	Code    ItemCode
}
