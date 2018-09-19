package main

import (
	"fmt"
	"sort"
)

var maxVal = 0

type Launch struct {
	attack  []int8
	defense []int8
}

type ResultCounter struct {
	attackCounter   int
	defenseCounter  int
	equalityCounter int
}

func DefaultCounter() ResultCounter {
	val := new(ResultCounter)
	val.attackCounter = 0
	val.defenseCounter = 0
	val.equalityCounter = 0
	return *val
}

func main() {
	seedRandom()

	numberOfLaunches := 1000000
	chanLaunch := make(chan Launch)
	go getLaunches(numberOfLaunches, chanLaunch)

	finalResult := DefaultCounter()

	for i := 0; i < numberOfLaunches; i++ {
		element := <-chanLaunch
		element.AppendResult(&finalResult)
	}

	finalResult.printResult(numberOfLaunches)
}

func (l Launch) AppendResult(result *ResultCounter) {
	winner := l.WhoIsWinning()
	if winner == ATTACK {
		result.attackCounter++
	} else if winner == DEFENSE {
		result.defenseCounter++
	} else if winner == EQUALITY {
		result.equalityCounter++
	} else {
		panic("wrong test value")
	}
}

func (l Launch) WhoIsWinning() int8 {
	sort.Sort(byReverseVal(l.attack))
	sort.Sort(byReverseVal(l.defense))

	attackWin := 0
	defenseWin := 0

	if l.attack[0] > l.defense[0] {
		attackWin++
	} else {
		defenseWin++
	}
	if l.attack[1] > l.defense[1] {
		attackWin++
	} else {
		defenseWin++
	}

	if attackWin == 2 {
		return ATTACK
	} else if defenseWin == 2 {
		return DEFENSE
	} else {
		return EQUALITY
	}
}

func (finalResult ResultCounter) printResult(numberOfLaunches int) {
	fmt.Printf("%+v\n", finalResult)
	fmt.Printf("number of launch: %v\n", numberOfLaunches)
	percentAttack := (float64(finalResult.attackCounter) / float64(numberOfLaunches)) * 100
	fmt.Printf("percent attack win: %v%%\n", percentAttack)

	percentDefense := (float64(finalResult.defenseCounter) / float64(numberOfLaunches)) * 100
	fmt.Printf("percent defense win: %v%%\n", percentDefense)

	percentEquality := (float64(finalResult.equalityCounter) / float64(numberOfLaunches)) * 100
	fmt.Printf("percent attack win: %v%%\n", percentEquality)

}
