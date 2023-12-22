package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"time"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var reaminingTickets uint = conferenceTickets
	var bookings = make([]map[string]string, 0) // empty list of maps

	greetUsers(conferenceName, conferenceTickets, reaminingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, reaminingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookings, reaminingTickets := bookTickets(&reaminingTickets, &userTickets, &bookings, firstName, lastName, email, conferenceName)

			go sendTicket(userTickets, firstName, lastName, email) // introducing a go keyword handles to concurrency

			userNames := getUserNames(bookings)
			fmt.Printf("The usernames of bookings are : %v\n", userNames)

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

func getUserNames(bookings []map[string]string) []string {
	userNames := []string{}
	for _, booking := range bookings {
		var fname string = booking["firstName"]
		var lname string = booking["lastName"]
		if len(userNames) != 0 {
			userNames = append(userNames, "\b, "+fname+" "+lname)
		} else {
			userNames = append(userNames, fname+" "+lname)
		}
	}
	return userNames
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

func bookTickets(reaminingTickets *uint, userTickets *uint, bookings *[]map[string]string, firstName string, lastName string, email string, conferenceName string) ([]map[string]string, uint) {

	// create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(*userTickets), 10)

	*reaminingTickets = *reaminingTickets - *userTickets
	*bookings = append(*bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get the confirmation at %v\n", firstName, lastName, *userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", *reaminingTickets, conferenceName)

	return *bookings, *reaminingTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v to \nemail address %v\n", ticket, email)
	fmt.Println("##############")
}
