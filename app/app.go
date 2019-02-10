package main

import (
	"fmt"

	"github.com/worming004/riskdice/logic"
)

func main() {
	logic.SeedRandom()

	logic.SetConfiguration()

	configurationSglt := logic.GetConfiguration()

	switch configurationSglt.Mode {
	case logic.MINFORPERCENT:

		logic.MinForPercentMode()
		break
	case logic.PROBABILITY:

		logic.ProbabilityMode()
		break
	case logic.SINGLE:

		logic.SingleMode()
		break
	default:
		fmt.Println("No mode found")
		break
	}
}
