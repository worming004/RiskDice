package main

import (
	"fmt"
)

const DEFAULT_NUMBER_SIMULATION int = 10000

func main() {
	seedRandom()

	SetConfiguration()

	configurationSglt := GetConfiguration()

	switch configurationSglt.Mode {
	case MINFORPERCENT:

		MinForPercentMode()
		break
	case PROBABILITY:

		ProbabilityMode()
		break
	case SINGLE:

		SingleMode()
		break
	default:
		fmt.Println("No mode found")
		break
	}
}
