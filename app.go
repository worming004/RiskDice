package main

import (
	"flag"
	"fmt"
	"strings"
)

const DEFAULT_NUMBER_SIMULATION int = 10000

func main() {
	seedRandom()

	appModePtr := flag.String("mode", "MinForPercent",
		"Launch the app in the way to search how much attacking units is necessary to invade a territory")

	flag.Parse()

	switch strings.ToLower(*appModePtr) {
	case strings.ToLower("MinForPercent"),
		strings.ToLower("MFP"):

		MinForPercentMode()
		break
	case strings.ToLower("Probability"),
		strings.ToLower("Prob"),
		strings.ToLower("P"):

		ProbabilityMode()
	default:
		fmt.Println("No mode found")
		break
	}
}
