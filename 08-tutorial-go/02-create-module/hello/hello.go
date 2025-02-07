package main

import (
	"fmt"

	"cosmocarpincho.com/greetings"
)

func main() {
	message := greetings.Hello("Cosmo")
	fmt.Println(message)
}
