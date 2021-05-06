package test

import (
	"fmt"
	"github.com/shopspring/decimal"
	"go_test/common"
)

func Test_method() {
	fmt.Printf("in the %v.\n", common.GetFuncName())

	var n1 float64 = 3.1
	var n2 float64 = 4.1
	d1 := decimal.NewFromFloat(n1).Add(decimal.NewFromFloat(n2))
	fmt.Println(n1 + n2)
	fmt.Println(d1)
}
