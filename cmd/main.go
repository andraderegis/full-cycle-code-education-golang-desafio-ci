package main

import (
	"fmt"

	arithmetics "github.com/andraderegis/golang-fullcycle-desafio-ci/arithmetics"
)

func main() {
	var sum int = arithmetics.Sum(5, 5)

	fmt.Printf("The value sum of 5 + 5 is: %v\n", sum)
}
