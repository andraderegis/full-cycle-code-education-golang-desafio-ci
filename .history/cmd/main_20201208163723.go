package main

import (
	"fmt"
	arithmetics "golang-fullcycle-desafio-ci/arithmetics"
	"time"
)

// func main() {
// 	var sum int = arithmetics.Sum(5, 5)

// 	fmt.Printf("The value sum of 5 + 5 is: %v\n", sum)

// 	fmt.Println("To exit, type 'exit'")
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("=> ")
// 		text, _ := reader.ReadString('\n')
// 		text = strings.Replace(text, "\n", "", -1)

// 		switch text {
// 		case "exit":
// 			fmt.Println("Exiting...")
// 			os.Exit(0)
// 		}

// 		time.Sleep(1000 * time.Millisecond)
// 	}
// }

func main() {
	var sum int = arithmetics.Sum(5, 5)

	fmt.Printf("The value sum of 5 + 5 is: %v\n", sum)

	for {
		fmt.Println(time.Now().UTC())
		time.Sleep(1000 * time.Millisecond)
	}
}
