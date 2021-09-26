package main

import (
	"fmt"
)

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

// indianTax structはcalculateTax()をメンバメソッドにもつので、
// taxSystem interfaceを実装した、といえる。
func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

// singaporeTax structはcalculateTax()をメンバメソッドにもつので、
// taxSystem interfaceを実装した、といえる。
func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func main() {
	/*
		var system taxSystem

		system = &indianTax{taxPercentage: 30, income: 1000}
		tax := system.calculateTax()
		fmt.Println(tax)

		system = &singaporeTax{taxPercentage: 10, income: 2000}
		tax = system.calculateTax()
		fmt.Println(tax)
	*/

	calculators := []taxSystem{&indianTax{taxPercentage: 30, income: 1000}, &singaporeTax{taxPercentage: 10, income: 2000}}
	for _, calculator := range calculators {
		tax := calculator.calculateTax()
		fmt.Println(tax)
	}
}
