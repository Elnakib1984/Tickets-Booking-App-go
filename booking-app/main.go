package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// making a shortcut for Println function
var pl = fmt.Println
var pf = fmt.Printf
var s = fmt.Scan
var spf = fmt.Sprintf

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	pf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T. \n", conferenceTickets, remainingTickets, conferenceName)

	pf("Welcome to %v booking application\n", conferenceName)
	pf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	pl("Get your tickets here to attend")

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)

		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		pf("The first names of bookings are: %v\n", firstNames)

		var noTicketsRemaining bool = remainingTickets == 0
		if noTicketsRemaining {
			// end program
			pl("Our conference is booked out. come back next year.")

			//break
		}

	} else {

		if !isValidName {
			pl("You might have written a short first name or last name")

		}
		if !isValidEmail {
			pl("Please check that your email is written correctly")
		}
		if !isValidTicketNumber {
			pl("Number of tickets you entered is invalid")
		}

	}
	wg.Wait()

}

func greetUsers() {
	pf("Welcome to %v booking application \n", conferenceName)
	pf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	pl("Get your teickets here to attend")
}

func getFirstNames() []string { //within brackets we have the input parameters
	//and outside those brackets we have the output parameters

	//bookings is a slice with datatype string

	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask the user for their name
	pl("Enter your first name : ")
	s(&firstName)

	pl("Enter your last name : ")
	s(&lastName)

	pl("Enter your email address : ")
	s(&email)

	//pl(&remainingTickets)
	// if we write & before the variable we get the memory location
	pl("Enter number of tickets : ")
	s(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	pf("List of Bookings is %v\n", bookings)

	pf("Thank you %v %v for booking %v tickets from our application. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	pf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firsrName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = spf("%v tickets for %v %v", userTickets, firsrName, lastName)
	pl("##########################################")
	pf("Sending ticket:\n %v \n to email address %v \n", ticket, email)
	pl("##########################################")
	wg.Done()

}
