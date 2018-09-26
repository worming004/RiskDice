package main

const ATTACKER_WINNER int = 1
const DEFENDER_WINNER int = 2

type MultipleSimulationResult struct {
	attackIsWinner     int
	defenseIsWinner    int
	numberOfSimulation int
}

func MultipleSimulate(numberOfAttack, numberOfDefense int, numberOfSimulation int) MultipleSimulationResult {
	result := new(MultipleSimulationResult)
	result.numberOfSimulation = numberOfSimulation

	for i := 0; i < numberOfSimulation; i++ {
		singleResult := Simulate(numberOfAttack, numberOfDefense)
		if singleResult == ATTACKER_WINNER {
			result.attackIsWinner++
		} else if singleResult == DEFENDER_WINNER {
			result.defenseIsWinner++
		} else {
			panic("expected winner")
		}

	}

	return *result
}

func Simulate(numberOfAttack, numberOfDefense int) int {
	for numberOfAttack > 0 && numberOfDefense > 0 {
		launch := GenerateLaunch(int8(min(3, numberOfAttack)), int8(min(2, numberOfDefense)))
		result := launch.WhoIsWinning()
		numberOfAttack = numberOfAttack - int(result.attackLost)
		numberOfDefense = numberOfDefense - int(result.defenseLost)
	}
	if numberOfAttack > 0 {
		return ATTACKER_WINNER
	} else if numberOfDefense > 0 {
		return DEFENDER_WINNER
	} else {
		panic("winner expected")
	}
}
