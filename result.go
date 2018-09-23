package main

import (
	"fmt"
)

func (l *Launch) WhoIsWinning() LaunchResult {
	if l.result != nil {
		return *l.result
	}

	result := new(LaunchResult)

	for i := 0; i < min(len(l.attack), len(l.defense)); i++ {
		if l.attack[i] > l.defense[i] {
			result.defenseLost++
		} else {
			result.attackLost++
		}
	}
	l.result = result
	return *result
}

func (sim MultipleSimulationResult) PrintResult() {
	fmt.Printf("number of launches: %d\n", sim.numberOfSimulation)
	fmt.Printf("number of attacker winner: %d\n", sim.attackIsWinner)
	fmt.Printf("%% of attacker winner: %v\n", 100*sim.attackIsWinner/sim.numberOfSimulation)
	fmt.Printf("number of defenser winner: %d\n", sim.defenseIsWinner)
	fmt.Printf("%% of defenser winner: %v\n", 100*sim.defenseIsWinner/sim.numberOfSimulation)
}
