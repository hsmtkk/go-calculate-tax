package tax

type TaxSystem interface {
	CalculateTax() int
}

type IndianTax struct {
	TaxPercentage int
	Income        int
}

// IndianTax structはCalculateTax()をメンバメソッドにもつので、
// taxSystem interfaceを実装した、といえる。
func (i *IndianTax) CalculateTax() int {
	tax := i.Income * i.TaxPercentage / 100
	return tax
}

type SingaporeTax struct {
	TaxPercentage int
	Income        int
}

// SingaporeTax structはCalculateTax()をメンバメソッドにもつので、
// taxSystem interfaceを実装した、といえる。
func (i *SingaporeTax) CalculateTax() int {
	tax := i.Income * i.TaxPercentage / 100
	return tax
}
