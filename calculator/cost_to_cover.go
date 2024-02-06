package calculator

// CostToCover receives the Distance to be covered (D), the Distance covered Per Liter (DPL)
// and the Cost Per Liter of the autombile's fuel (CPL). With it, it calculates how much it will Cost To Cover
// the distance (CTC) following this formula: CTC = (D / DPL) * CPL
func CostToCover(
	totalDistance float64,
	distancePerLiter float64,
	costPerLiter float64,
) float64 {
	return (totalDistance / distancePerLiter) * costPerLiter
}
