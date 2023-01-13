package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if isUnitINUnits(units, unit) {
		bill[item] += units[unit]
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	newQuantity := bill[item] - units[unit]

	if !isItemINBill(bill, item) || !isUnitINUnits(units, unit) || newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = bill[item] - units[unit]
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if !isItemINBill(bill, item) {
		return 0, false
	}

	return bill[item], true
}

func isItemINBill(bill map[string]int, item string) bool {
	_, ok := bill[item]
	return ok
}

func isUnitINUnits(units map[string]int, unit string) bool {
	_, ok := units[unit]
	return ok
}
