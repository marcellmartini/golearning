package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var ir float32
	switch {
	case balance < 0:
		ir = 3.213
	case balance >= 0 && balance < 1000:
		ir = 0.5
	case balance >= 1000 && balance < 5000:
		ir = 1.621
	case balance >= 5000:
		ir = 2.475
	}
	return ir
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	for years = 0; balance < targetBalance; years++ {
		balance = AnnualBalanceUpdate(balance)
	}
	return years
}
