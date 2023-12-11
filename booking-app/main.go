package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var reaminingTickets uint = conferenceTickets
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, reaminingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, reaminingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookings, reaminingTickets := bookTickets(&reaminingTickets, &userTickets, &bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if reaminingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("FirstName and LastName must contain at least 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Email must contain '@'")
			}
			if !isValidUserTickets {
				fmt.Printf("Available number of tickets : %v. Enter data again to get these many tickets.\n", reaminingTickets)
			} else {
				fmt.Println("Invalid input. Please check again")
			}
		}
	}

}

func greetUsers(confName string, confTickets int, remTickets uint) {
	fmt.Printf("Welcome to %v booking application.\n", confName)
	fmt.Println("We have a total of", confTickets, "and", remTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{} // collect firstNames to this slice
	// In Go, '_' are used to identify unused variables
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(reaminingTickets *uint, userTickets *uint, bookings *[]string, firstName string, lastName string, email string, conferenceName string) ([]string, uint) {
	*reaminingTickets = *reaminingTickets - *userTickets
	*bookings = append(*bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get the confirmation at %v\n", firstName, lastName, *userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", *reaminingTickets, conferenceName)

	return *bookings, *reaminingTickets
}
