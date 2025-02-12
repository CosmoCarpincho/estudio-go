package main

import (
	"fmt"
	"math"
	"runtime"
  "time"

)

func main_0() {
	fmt.Println("Hola")

}

// FOR
func main_for() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

}

// for continuo
func main_for_continuo() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// se puede quitar los  ;
func main_for_continuo_sin() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// bucle infinito
func bucle_infinito() {
	for {
		fmt.Println("infinito")
	}
}

// if

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}

func main_if() {
	fmt.Println(sqrt(2), sqrt(-4))
}

// if con breve declaración que se ejecuta antes que la condiciónew
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main_if_con_declaracion() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// if else

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)

	}

	return lim
}

func main_if_else() {
  fmt.Println(
    pow2(3, 2, 10),
    pow2(3, 3, 20),
  )
}


// switch



func main_swtich() {
  fmt.Print("Go se ejecuta en ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd
    // plan9, windows...
    fmt.Printf("%s. \n", os)
  }
}


// swtich 2
func main_swtich_2() {
  fmt.Println("¿Cuándo es sábado?")
  today := time.Now().Weekday()
  println("time.Saturday = ", time.Saturday)
  switch time.Saturday {
  case today + 0:
    fmt.Println("Hoy.")
  case today + 1:
    fmt.Println("Mañana.")
  case today + 2:
    fmt.Println("En dos días")
  default:
    fmt.Println("Falta mucho.")
  }
}

// switch sin condición (forma limpia de hacer if then else)
func main_switch() {
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Buenos días!")
  case t.Hour() < 17:
    fmt.Println("Buenas tardes.")
  default:
    fmt.Println("Buenas noches.")
  }
}


// Defer
func main_defer() {
  defer fmt.Println("mundo")

  fmt.Println("hola")
}

func main_pila_defer() {
  fmt.Println("contando")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("listo")
}


// punteros

func main_puteros() {
  i, j := 42, 2701

  p := &i
  fmt.Println(*p)
  *p = 21
  fmt.Println(i)

  p = &j
  *p = *p / 37
  fmt.Println(j)
}


type Vertex struct {
  X int
  Y int
}

func main_struct() {
  fmt.Println(Vertex{1, 2})
}

// struct

func main_struct_2() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)
}


// punteros a structs
func main_punteros_a_structs() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}

var (
  v1 = Vertex{1, 22}
  v2 = Vertex{X: 1}
  v3 = Vertex{}
  p = &Vertex{1, 2}
)

func main() {
  fmt.Println(v1, p, v2, v3)
}
