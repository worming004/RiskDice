package main

import (
	"math/rand"
	"sort"
	"time"
)

type Launch struct {
	attack  []int8
	defense []int8
	result  *LaunchResult
}

type LaunchResult struct {
	attackLost  int8
	defenseLost int8
}

func seedRandom() {
	rand.Seed(time.Now().UTC().Unix())
}

func getLaunches(numberOfItem int, launchChan chan Launch, numAttackDices, numDefendeDices int) {
	for i := 0; i < numberOfItem; i++ {
		launchChan <- generateLaunch(3, 2)
	}
}

func generateLaunch(attacks, defenders int8) Launch {
	numberOfDice := attacks + defenders
	token5 := make([]int8, numberOfDice)
	for i := int8(0); i < numberOfDice; i++ {
		token5[i] = getDiceValue()
	}

	launch := new(Launch)
	launch.attack = token5[0:attacks]
	launch.defense = token5[attacks:numberOfDice]

	sort.Sort(byReverseVal(launch.attack))
	sort.Sort(byReverseVal(launch.defense))

	return *launch
}

func getDiceValue() int8 {
	return int8(rand.Intn(6))
}
