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
	fmt.Printf("%% of attacker winner: %v\n", sim.GetAttackWinnerPercent())
	fmt.Printf("number of defenser winner: %d\n", sim.defenseIsWinner)
	fmt.Printf("%% of defenser winner: %v\n", sim.GetDefenseWinnerPercent())
}

func (sim MultipleSimulationResult) GetAttackWinnerPercent() float32 {
	return float32(100 * sim.attackIsWinner / sim.numberOfSimulation)
}

func (sim MultipleSimulationResult) GetDefenseWinnerPercent() float32 {
	return float32(100 * sim.defenseIsWinner / sim.numberOfSimulation)
}
