package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var reaminingTickets uint = conferenceTickets
	var bookings []string

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and", reaminingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Printf("Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter the number of tickets you want to buy : ")
		fmt.Scan(&userTickets)

		reaminingTickets = reaminingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will get the confirmation at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", reaminingTickets, conferenceName)

		firstNames := []string{} // collect firstNames to this slice
		// In Go, '_' are used to identify unused variables
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are : %v\n", firstNames)
	}

}
