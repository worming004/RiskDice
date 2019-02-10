package main

func FindMinimumAttackToWin(numberOfDefenser int, expectedPercent float32) (int, float32) {
	if expectedPercent < 0 || expectedPercent > 100 {
		panic("selected percentage is not between 0 and 100")
	}
	for numberOfAttacker := 1; true; numberOfAttacker++ {
		nbrSmimulation := GetConfiguration().NbrOfSimulation
		result := MultipleSimulate(numberOfAttacker, numberOfDefenser, nbrSmimulation)
		lastPercentFound := result.GetAttackWinnerPercent()
		if lastPercentFound >= expectedPercent {
			return numberOfAttacker, lastPercentFound
		}
	}
	panic("")
}
