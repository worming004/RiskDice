package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
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

func MinForPercentMode() {
	err, defenders, expectedPercent := GetDefendersPercent(flag.Args())
	if err != nil {
		fmt.Println("error while parsing arguments.\n this program will use default values")
		fmt.Println("call '<appname> -<mode> <number of defender> <expected winrate>' to use this app")
		defenders, expectedPercent = 15, 80
	}

	unitsToWin := FindMinnimumAttackToWin(defenders, expectedPercent)
	fmt.Printf("minimum of %v unit is necessary to beat %v with a %v winrate\n", unitsToWin, defenders, expectedPercent)
	// result := MultipleSimulate(15, 15, DEFAULT_NUMBER_SIMULATION)
	// result.PrintResult()
}

func GetDefendersPercent(args []string) (error, int, float32) {
	if len(args) != 2 {
		return errors.New("wrong arguments"), 0, 0
	}
	defenders, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return err, 0, 0
	}
	expectedPercent, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		return err, 0, 0
	}
	return nil, int(defenders), float32(expectedPercent)
}

func ProbabilityMode() {
	err, attackers, defenders := GetAttackersDefenders(flag.Args())
	if err != nil {
		panic("Error in ProbabilityMode")
	}
	prop := MultipleSimulate(attackers, defenders, DEFAULT_NUMBER_SIMULATION)
	prop.PrintResult()
}

func GetAttackersDefenders(args []string) (error, int, int) {
	if len(args) != 2 {
		return errors.New("wrong arguments"), 0, 0
	}
	attackers, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return err, 0, 0
	}
	defenders, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		return err, 0, 0
	}
	return nil, int(attackers), int(defenders)
}
