package main

import (
	"fmt"
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

func GenerateLaunch(attacks, defenders int8) Launch {
	if attacks > 3 {
		fmt.Println("attacks > 3 for a dice battle. Are you sure you do what you wants ?")
	}
	if defenders > 2 {
		fmt.Println("defenders > 2 for a dice battle. Are you sure you do what you wants ?")
	}
	numberOfDice := attacks + defenders
	token5 := make([]int8, numberOfDice)
	for i := int8(0); i < numberOfDice; i++ {
		token5[i] = getRandomDiceValue()
	}

	launch := new(Launch)
	launch.attack = token5[0:attacks]
	launch.defense = token5[attacks:numberOfDice]

	sort.Sort(byReverseVal(launch.attack))
	sort.Sort(byReverseVal(launch.defense))

	return *launch
}

func getRandomDiceValue() int8 {
	return int8(rand.Intn(6))
}

func (lr LaunchResult) String() string {
	return fmt.Sprintf("attacker lost %v unit(s) and defender lost %v unit(s)", lr.attackLost, lr.defenseLost)
}
