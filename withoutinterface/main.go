package main

import (
	"fmt"
)

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func main() {
	indianSystem := &indianTax{taxPercentage: 30, income: 1000}
	tax := indianSystem.calculateTax()
	fmt.Println(tax)

	singaporeSystem := &singaporeTax{taxPercentage: 10, income: 2000}
	tax = singaporeSystem.calculateTax()
	fmt.Println(tax)
}
