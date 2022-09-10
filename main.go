package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUsers()

	for {
		// get user input
		firstName,lastname,email,userTickets := getUserInput()

		// validate user information
		isValidName,isValidEmail,isValidTicketNumber := validUserInput(firstName, lastname, email, userTickets, remainingTickets)


		if isValidName && isValidEmail && isValidTicketNumber {
			// user booking the tickets
			bookTicket(remainingTickets,userTickets,bookings,firstName,lastname,email,conferenceName)

			// print first names
			firstNames := FirstNames()
			fmt.Printf("The first names of of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last namer you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println("Your input data is invalid, Try Again!")
		}
	}
}

func getFirstNames(bookings []string) {
	panic("unimplemented")
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}


func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validUserInput(firstName string,lastname string, email string,userTickets uint, remainingTickets uint)(bool,bool,bool){
	isValidName := len(firstName) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput()(string,string,string,uint){
	var firstName string
	var lastname string
	var email string
	var userTickets uint
	// getting input from the user
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName,lastname,email,userTickets
}

func bookTicket(remainingTickets uint,userTickets uint,bookings []string,firstName string, lastname string,conferenceName string,email string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastname)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}