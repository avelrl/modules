package module_main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func DoingSomeInteresting() string {
	d, err := decimal.NewFromString("1.00001")
	if err != nil {
		return "0.0"
	}
	fmt.Printf("%s\n", d)
	return d.String()
}
