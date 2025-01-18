package main

import (
	"fmt"
	"log"
	"net/http"
)

func main_1() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

// panic
func main_2() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

}

// panic
func main_3() {
	fmt.Println("Start")
	panic("something bad happened")
	fmt.Println("end")
}

// si queremos usar dos veces el mismo puerto (2 veces go run ...)
// generamos painc -> err.Error() . Que no se ejecute el programa
func main_4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func main_5() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
	// OJO VER SALIDA ->
	// -> start
	// -> this was deferred
	// -> panic: something bad happened

	// Los defer van a ejecutarce aunque entre en panic.(antes del panic)

}

func main_6() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")

	// ➜ go run .\main2.go
	// start
	// about to panic
	// 2025/01/17 21:57:45 Erro: something bad happened
	// end
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Erro:", err)
		}
	}()

	panic("something bad happened")
	fmt.Println("done panicking")
}
