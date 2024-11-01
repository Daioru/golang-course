package main

func getMonthlyPrice(tier string) int {
	var price int
	switch tier {
	case "basic":
		price = 10000
	case "premium":
		price = 15000
	case "enterprise":
		price = 50000
	default:
		price = 0
	}
	return price
}
