package main

import(
  "fmt"
)

func main() {
  conferenceName := "Go Conference"
  const conferenceTickets = 50
  var remainingTickets uint = 50

  fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

  fmt.Printf("Welcome to %v booking application\n", conferenceName) // Se agrega espacio antes y despues de la variable automaticamente
  fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
  fmt.Println("Get yout tickets here to attend")

  var userName string
  var userTickets int

  fmt.Scan(&userName)

  userName = "Tom"
  userTickets = 2

  fmt.Printf("Use %v booked %v tickets.\n", userName, userTickets)
}

// 40:59
