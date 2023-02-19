package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRateDecimal := float64(successRate) / 100.0
    workingCarsPerHour := float64(productionRate) * successRateDecimal
    return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	groupSize := 10
	groupCost := 95000
	individualCost := 10000

	numGroups := carsCount / groupSize
	numIndividuals := carsCount % groupSize

	totalCost := numGroups*groupCost + numIndividuals*individualCost
	return uint(totalCost)
}
