package logic

import (
	"errors"
	"flag"
	"strconv"
)

func ProbabilityMode() {
	err, attackers, defenders := GetAttackersDefendersArgs(flag.Args())
	if err != nil {
		panic("Error in ProbabilityMode")
	}
	prop := MultipleSimulate(attackers, defenders, GetConfiguration().NbrOfSimulation)
	prop.PrintResult()
}

func GetAttackersDefendersArgs(args []string) (error, int, int) {
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
