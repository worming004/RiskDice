package main

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
