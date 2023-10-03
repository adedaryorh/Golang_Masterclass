package main

import "fmt"

func main() {
	fmt.Println("Hello world !!")
	app()
}

func app() {

	conferenceName := "Golang Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to the %v booking app\n", conferenceName)
	fmt.Printf("we have total no of %v tickets and %v are still available.\n ", conferenceTicket, remainingTickets)
	fmt.Println("Kindly supply your name book your tickets")

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your firstName")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email Name")
	fmt.Scan(&email)
	fmt.Println("Enter Number of tickets: ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket

	fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTicket)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation email %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("remaing tickest is now %v for the %v", remainingTickets, conferenceName)

}
