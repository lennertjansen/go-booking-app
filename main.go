package main

import "fmt"

func main() {
    conferenceName := "Go Conference"
    const conferenceTickets int = 50
    var remainingTickets uint = 50
    var bookings []string

    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets left. And %v tickets are still available\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")

    var firstName string
    var lastName string
    var email string
    var userTickets uint

    fmt.Printf("Enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Printf("Enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Printf("Enter your email address: ")
    fmt.Scan(&email)

    fmt.Printf("Enter desired number of tickets: ")
    fmt.Scan(&userTickets)

    remainingTickets = remainingTickets - userTickets
    bookings = append(bookings, firstName + " " + lastName)

    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

    fmt.Printf("These are all our bookings: %v\n", bookings)
}
