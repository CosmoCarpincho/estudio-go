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

   type empleado struct {
     nombre string
     apellido string
   }

   var alguien = empleado { "cosmo", "programador"}

   fmt.Printf("Hola %v, trabajas de %v\n", alguien.nombre, alguien.apellido)


   type punto struct {
     x int
     y int
   }

   
   punto1 := punto {1, 10}
   punto2 := punto {10, 1}

   fmt.Println(punto1, punto2)
   fmt.Println(punto1.x)

   punteroStruct := &punto1

   fmt.Printf("puntero: %p valor x: %v valor y: %v\n", punteroStruct, punteroStruct.x, (*punteroStruct).y)
}
