package main

import (
	"math/rand"
	"time"
)

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

func seedRandom() {
	rand.Seed(time.Now().UTC().Unix())
}

func getLaunches(numberOfItem int, launchChan chan Launch) {
	for i := 0; i < numberOfItem; i++ {
		launchChan <- generateLaunch()
	}
}

func generateLaunch() Launch {
	token5 := make([]int8, 5)
	for i := 0; i < 5; i++ {
		token5[i] = getDiceValue()
	}

	launch := new(Launch)
	launch.attack = token5[0:3]
	launch.defense = token5[3:5]

	return *launch
}

func getDiceValue() int8 {
	return int8(rand.Intn(6))
}
