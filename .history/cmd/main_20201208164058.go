package main

import (
	"fmt"
	arithmetics "golang-fullcycle-desafio-ci/arithmetics"
	"time"
)

func main() {
	var sum int = arithmetics.Sum(5, 5)

	fmt.Printf("The value sum of 5 + 5 is: %v\n", sum)

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
