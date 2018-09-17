package main

import (
	"math/rand"
	"time"
)

func seedRandom() {
	rand.Seed(time.Now().UTC().Unix())
}

func getLaunches(numberOfItem int) []Launch {
	result := make([]Launch, numberOfItem)
	for i := 0; i < numberOfItem; i++ {
		result[i] = generateLaunch()
	}
	return result
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
	return int8(rand.Intn(7))
}
