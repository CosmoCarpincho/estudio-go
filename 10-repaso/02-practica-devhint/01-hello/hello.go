package main

import "fmt"


func main() {
  message := greetMe("World")
  fmt.Println(message)
}

func greetMe(name string) string {
  return "Hello, " + name + "!"
}
