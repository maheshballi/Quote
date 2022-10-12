package main

import (
	"booking_app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var booking [50]string  ----->Using array
// Using map
// var bookings = make([]map[string]string, 0)

// using struct
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickets {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//booking[0] = firstName + " " + lastName ---(Array)

		//fmt.Println("The whole array:", booking)
		//fmt.Println("The first value:", booking[0])
		//fmt.Printf("Array type: %T\n", booking)
		//fmt.Println("Arraylength:", len(booking))

		//Using slice

		//fmt.Println("The whole slice:", booking)
		//fmt.Println("The first value:", booking[0])
		//fmt.Printf("Slice Type: %T\n", booking)
		//fmt.Println("Slice length:", len(booking))

		// printing firstname
		firstName := getFirstNames()
		fmt.Printf("The first name of bookings are: %v\n", firstName)

		if remainingTickets == 0 {

			fmt.Println("Sorry all the tickets for our conference has been booked. come back next year!")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("The first name or lastname you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("You may not have entered @ sign, kindly enter it!")
		}
		if !isValidUserTickets {
			fmt.Println("You have entered invalid tickets")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and still %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend the conference")
}

func getFirstNames() []string {
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

	fmt.Println("Enter your first Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//using map
	//var userData = make(map[string]string)
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %s %s you have booked %v tickets, soon you will  recieve a confirmation at your mail %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(time.Second * 10)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################")

	fmt.Printf("sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()

}
