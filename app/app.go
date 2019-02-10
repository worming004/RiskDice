package main

import (
	"fmt"
)

func main() {
	SeedRandom()

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
