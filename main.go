package main

import (
	"fmt"
	utils "golang-fullcycle-desafio-ci/utils"
)

func main() {
	var sum int = utils.Sum(5, 5)

	fmt.Printf("The value sum of 5 + 5 is: %v\n", sum)
}
