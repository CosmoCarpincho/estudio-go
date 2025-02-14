// Pseudocodigo wikipedia: https://es.wikipedia.org/wiki/Test_de_primalidad_de_Miller-Rabin

// def test_Miller_Rabin(n: int, k: int) -> bool:
//     s, d = satisfacen n = 2**s * d + 1, d impar
//     repetir k veces:
//         a = entero al azar entre 2 y n-1
//         fpp = False # fuertemente primo en base a
//         if 1 == a**d % n:
//             fpp = True
//         else:
//             r = 0
//             while r  <= s and fpp == False:
//                 if n - 1 == a**(2**r * d) % n:
//                     fpp = True
//          if fpp == False: # si no pasa la prueba
//             return False
//     # si pasó las k pruebas
//     return True

package main

func main() {

}
