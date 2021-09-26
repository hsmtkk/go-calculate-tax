package tax_test

import (
	"testing"

	"github.com/hsmtkk/go-calculate-tax/tax"
	"github.com/stretchr/testify/assert"
)

func TestIndianTax(t *testing.T) {
	want := 300
	system := &tax.IndianTax{Income: 1000, TaxPercentage: 30}
	got := system.CalculateTax()
	assert.Equal(t, want, got)
}
