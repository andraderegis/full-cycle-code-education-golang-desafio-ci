package main

import (
	"bufio"
	"fmt"
	arithmetics "golang-fullcycle-desafio-ci/arithmetics"
	"os"
)

func main() {
	var sum int = arithmetics.Sum(5, 5)

	fmt.Printf("The value sum of 5 + 5 is: %v\n", sum)

	fmt.Println("To exit, type 'exit'")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')

		// text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "exit":
			os.Exit(0)
		}
	}
}
