package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

// package level variables cannot be defined using :=
var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// fixed array of strings: var bookings [50]string
// var bookings = []string{} // slices allow for dynamic inserts or var bookings []string
var bookings = make([]UserData, 0) //initialize a list/slice of maps, 0 is initial size

// allows multiple types and custom types
type UserData struct {
	firstName string
	lastName	string
	email			string
	userTickets		uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	// for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(userTickets, firstName, lastName, userEmail)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, userEmail)
			firstNames := getFirstNames()

			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is sold out. Come back next year.")
				// break
			}
		} else {
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("First or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address does not include @")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets is invalid")
			}
			// continue
		}
	// }
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// %T - type
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings { // _ - used to identify unused variables
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint // positive numbers only
	// user input
	// pointer - points to the memory address of another variable
	// fmt.Println(remainingTickets)  // 50
	// fmt.Println(&remainingTickets) // 0xc00001a0d0

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint, fName string, lName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData {
		firstName: fName,
		lastName: lName,
		email: email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v\n", fName, lName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done() // removes thread from wait group
}
