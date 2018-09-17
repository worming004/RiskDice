package main

import (
	"fmt"
	"sort"
)

var maxVal = 0

type Launch struct {
	attack  []int
	defense []int
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
	launches := getLaunches(numberOfLaunches)

	resultOfWinners := make([]string, numberOfLaunches)

	for i, element := range launches {
		winner := element.WhoIsWinning()
		resultOfWinners[i] = winner
	}

	finalResult := DefaultCounter()

	for _, element := range resultOfWinners {
		if element == "attack" {
			finalResult.attackCounter++
		} else if element == "defense" {
			finalResult.defenseCounter++
		} else if element == "equality" {
			finalResult.equalityCounter++
		} else {
			panic("wrong test value")
		}
	}

	fmt.Printf("%+v\n", finalResult)
}

func (l Launch) WhoIsWinning() string {
	sort.Sort(sort.Reverse(sort.IntSlice(l.attack)))
	sort.Sort(sort.Reverse(sort.IntSlice(l.defense)))
	// sort.Ints(l.attack)
	// sort.Ints(l.defense)

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
		return "attack"
	} else if defenseWin == 2 {
		return "defense"
	} else {
		return "equality"
	}
}
