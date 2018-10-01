package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
)

func SingleMode() {
	attackers, defenders, err := GetSingleModeArgs(flag.Args())
	if err != nil {
		panic("error in single mode")
	}
	result := GenerateLaunch(attackers, defenders)
	fmt.Println(result.WhoIsWinning())
}

func GetSingleModeArgs(args []string) (int8, int8, error) {
	if len(args) != 2 {
		return 0, 0, errors.New("wrong arguments")
	}
	attackers, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	defenders, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return int8(attackers), int8(defenders), nil
}
