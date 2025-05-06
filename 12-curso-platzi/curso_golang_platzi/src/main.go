package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("+ %d +", i)
	}

	fmt.Println()

	for i := range 10 {
		fmt.Printf("- %d -", i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++

		if counterForever == 20 {
			break
		}
	}

}
