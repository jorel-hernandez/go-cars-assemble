package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	//panic("SuccessRate not implemented")
	var rateValue float64
	
	if speed == 0 { 
		rateValue = 0
	} else if speed >= 1 && speed <= 4 {
		rateValue = 1
	} else if speed >= 5 && speed <= 8 {
		rateValue = 0.9
	} else {
		rateValue = 0.77
	}
	
	return rateValue
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	//panic("CalculateProductionRatePerHour not implemented")
	return float64(speed) * 221 * SuccessRate(speed)
	
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	//panic("CalculateProductionRatePerMinute not implemented")
	return int(CalculateProductionRatePerHour(speed) / float64(60))
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	//panic("CalculateLimitedProductionRatePerHour not implemented")
	
	//ratePerHour := CalculateProductionRatePerHour(speed)
	var ratePerHour float64 = CalculateProductionRatePerHour(speed)
	
	if (ratePerHour > limit) { 
		ratePerHour = limit
	} 
	
	return ratePerHour
}
