package main

import (
  "fmt"
)

type Vertex2 struct {
  X, Y int
}

var (
  v12 = Vertex2{1, 2}
  v22 = Vertex2{X: 1}
  v32 = Vertex2{}
  p2 = &Vertex2{1, 2}
)

func main_22() {
  fmt.Println("struct literal")

  fmt.Println(v12, p2, v22, v32)
}
