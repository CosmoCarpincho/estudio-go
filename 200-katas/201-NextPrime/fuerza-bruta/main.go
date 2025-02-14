// Get the next prime number!

// You will get a numbern (>= 0) and your task is to find the next prime number.

// Make sure to optimize your code: there will numbers tested up to about 10^12.

// Examples
// 5   =>  7
// 12  =>  13

// Interesante pero No. Voy a usar fuerza bruta.
// Criba de Eratóstenes
 // 1.Primer paso: listar los números naturales comprendidos entre 2 hasta el número que se desee, en este caso, hasta el 20.
 // 2. Segundo paso: se toma el primer número no rayado ni marcado, como número primo.
 // 3. Tercer paso: se tachan todos los múltiplos del número que se acaba de indicar como primo.
 // 4. Cuarto paso: si el cuadrado del primer número que no ha sido rayado ni marcado es inferior a 20, entonces se repite el segundo paso. Si no, el algoritmo termina, y todos los enteros no tachados son declarados primos.


package main

import (
  "fmt"
)

func NextPrime(n int) int {
  var primo int = -1;
  var isPrimo bool = false;

  if n < 2 && n >= 0 {
    return 2
  }

  for i:= n+1; !isPrimo; i++ {
    for num := 2; num < i; num++ {
      if i % num == 0 {
        isPrimo = false
        break
      }

      isPrimo = true
      primo = i 

    }

  }
  return primo
}

func main () {
  fmt.Println(NextPrime(0))
  fmt.Println(NextPrime(2))
  fmt.Println(NextPrime(3))
  fmt.Println(NextPrime(13))
  fmt.Println(NextPrime(181))
  fmt.Println(NextPrime(911))

}
