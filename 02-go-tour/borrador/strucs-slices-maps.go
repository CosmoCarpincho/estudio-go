package main

import (
  "fmt"
)

func main_6() {
  fmt.Println("Hola")
}

func main() {
  // punteros
  var p *int; // guarda dirección de memoria que almacena un int

  println(p) // -> 0x0  como apunta a cero es nil. Quiere decir que no tiene asignado una posición de memoria valida
  fmt.Println(p) // -> nil .

  //*p = 10 // nil

   var a int = 10
   p = &a

   fmt.Printf("p = %v  *p = %v\n", p, *p)

   type algo struct {
     nombre string
     apellido string
   }

   var alguien = algo{ "cosmo", "programador"}

   fmt.Printf("Hola %v, trabajas de %v\n", alguien.nombre, alguien.apellido)
}
