package main

import (
  "fmt"

)

func main() {
  sample := "⌘Gopher"

  fmt.Println("Bytes:")
  for i := 0; i < len(sample); i++ {
    fmt.Printf("%x ", sample[i])
  }
  fmt.Println()

  fmt.Println("Runes:")

  runes := []rune(sample)
  for i:= 0; i < len(runes); i++ {
    fmt.Printf("%q (%U)\n", runes[i], runes[i])


  }

  fmt.Println("Con range:")

  for _, r:= range sample {
    fmt.Printf("%q (%U)\n", r, r)
  }


}
