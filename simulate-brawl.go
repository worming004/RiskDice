package main

const AttackWinner int = 1
const DefenseWinner int = 2

type MultipleSimulationResult struct {
	attackIsWinner     int
	defenseIsWinner    int
	numberOfSimulation int
}

func Simulate(numberOfAttack, numberOfDefense int) int {
	for numberOfAttack > 0 && numberOfDefense > 0 {
		launch := generateLaunch(int8(min(3, int(numberOfAttack))), int8(min(2, int(numberOfDefense))))
		result := launch.WhoIsWinning()
		numberOfAttack = numberOfAttack - int(result.attackLost)
		numberOfDefense = numberOfDefense - int(result.defenseLost)
	}
	if numberOfAttack > 0 {
		return AttackWinner
	} else if numberOfDefense > 0 {
		return DefenseWinner
	} else {
		panic("winner expected")
	}
}

func MultipleSimulate(numberOfAttack, numberOfDefense int, numberOfSimulation int) MultipleSimulationResult {
	result := new(MultipleSimulationResult)
	result.numberOfSimulation = numberOfSimulation

	for i := 0; i < numberOfSimulation; i++ {
		singleResult := Simulate(numberOfAttack, numberOfDefense)
		if singleResult == AttackWinner {
			result.attackIsWinner++
		} else if singleResult == DefenseWinner {
			result.defenseIsWinner++
		} else {
			panic("expected winner")
		}

	}

	return *result
}
