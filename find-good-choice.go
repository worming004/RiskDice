package main

func FindMinnimumAttackToWin(numberOfDefenser int, expectedPercent float32) int {
	for numberOfAttacker := 1; true; numberOfAttacker++ {
		result := MultipleSimulate(numberOfAttacker, numberOfDefenser, DEFAULT_NUMBER_SIMULATION)
		lastPercentFound := result.GetAttackWinnerPercent()
		if lastPercentFound >= expectedPercent {
			return numberOfAttacker
		}
	}
	panic("")
}
