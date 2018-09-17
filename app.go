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

	numberOfLaunches := 100000000
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
	fmt.Printf("number of launch: %v\n", numberOfLaunches)
	percentAttack := (float64(finalResult.attackCounter) / float64(numberOfLaunches)) * 100
	fmt.Printf("percent attack win: %v%%\n", percentAttack)

	percentDefense := (float64(finalResult.defenseCounter) / float64(numberOfLaunches)) * 100
	fmt.Printf("percent defense win: %v%%\n", percentDefense)

	percentEquality := (float64(finalResult.equalityCounter) / float64(numberOfLaunches)) * 100
	fmt.Printf("percent attack win: %v%%\n", percentEquality)

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
