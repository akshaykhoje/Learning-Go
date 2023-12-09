package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var reaminingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and", reaminingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter the number of tickets you want to buy : ")
	fmt.Scan(&userTickets)

	reaminingTickets = reaminingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get the confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", reaminingTickets, conferenceName)
}
