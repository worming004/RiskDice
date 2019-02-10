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
		"Launch the app in the way to search how much attacking units is necessary to invade a territory"+
			"\n MinForPercent \n Probability \n  Single")

	nbrSimul := flag.Int("nbrsimul", DEFAULT_NUMBER_SIMULATION,
		"Number of launch to do to execute the simulation. Utilisation : -nbrsimul <value>")

	showConfigurationPtr := flag.Bool("conf", false,
		"Show the configuration of the application. Utilisation: -conf=<true/false>")

	flag.Parse()
	showConfiguration := *showConfigurationPtr

	configurationSglt := GetConfiguration()
	configurationSglt.NbrOfSimulation = *nbrSimul

	switch strings.ToLower(*appModePtr) {
	case strings.ToLower("MinForPercent"),
		strings.ToLower("MFP"):
		configurationSglt.Mode = MINFORPERCENT
		break
	case strings.ToLower("Probability"),
		strings.ToLower("Prob"),
		strings.ToLower("P"):
		configurationSglt.Mode = PROBABILITY
		break
	case strings.ToLower("Single"),
		strings.ToLower("S"):
		configurationSglt.Mode = SINGLE
		break
	default:
		fmt.Println("No mode found")
		break
	}

	if showConfiguration {
		fmt.Println("Configuration:")
		fmt.Println(*configurationSglt)
	}

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
