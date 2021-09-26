package main

import (
	"fmt"

	"github.com/hsmtkk/go-calculate-tax/tax"
)

func main() {
	systems := []tax.TaxSystem{&tax.IndianTax{TaxPercentage: 30, Income: 1000}, &tax.SingaporeTax{TaxPercentage: 10, Income: 2000}}
	calculateTotalTax(systems)
}

func calculateTotalTax(taxSystems []tax.TaxSystem) {
	total := 0
	for _, system := range taxSystems {
		tax := system.CalculateTax()
		total += tax
	}
	fmt.Println(total)
}
