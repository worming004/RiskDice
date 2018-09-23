package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
)

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
