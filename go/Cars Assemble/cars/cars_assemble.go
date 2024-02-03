package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	workingCarsPerHour := float64(productionRate) * (successRate / 100)
	
	return float64(workingCarsPerHour)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	workingCarsPerMinute := int(workingCarsPerHour / 60)

	return workingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	decimalQnty := carsCount / 10
	fractionalQnty := carsCount - (decimalQnty * 10)

	costOfProduction := decimalQnty * 95000 + fractionalQnty * 10000

	return uint(costOfProduction)
}
