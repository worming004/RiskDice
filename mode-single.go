package main

import (
	"fmt"
)

func SingleMode() {
	result := GenerateLaunch(10, 10)
	fmt.Println(result.WhoIsWinning())
}
