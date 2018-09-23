package main

const DEFAULT_NUMBER_SIMULATION int = 100

func main() {
	seedRandom()

	result := MultipleSimulate(15, 15, DEFAULT_NUMBER_SIMULATION)
	result.PrintResult()
}
