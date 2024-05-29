package main

import (
	"fmt"
	"strings"
)

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func userInput[k comparable](what string, input k) k {
	fmt.Printf("Please enter your %v: ", what)
	fmt.Scan(&input)
	return input
}

func main() {
	// Var conferance
	conferanceName := "Go Conferance"
	const conferaneTickets = 10
	var remainingTickets uint = 10
	var bookings []string // <-- slice

	// Printing some welcome message to the user
	fmt.Printf("Welcome to our %v booking Application\n", conferanceName)
	fmt.Printf("we have total of %v tickets and %v are still remaining to take\n", conferaneTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n======================\n")

	// printing the TYPE of an VARIALBLE
	// fmt.Printf("%T\n", conferanceName)

	var userName string
	var userEmail string
	var userTickets uint
	// for {} <-- this kind of loop will never ends if you define like this
	for {
		userName = userInput("full name", userName)
		userEmail = userInput("email address", userEmail)
		fmt.Print("Please enter how many tickets you want? ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("You're booking out of range number of bookings: %v is remaining\n", remainingTickets)
			break
		}
		// userEmail= userInput[int]("ticket numbers", userTicket)
		bookings = append(bookings, userName+" "+userEmail)
		remainingTickets = remainingTickets - userTickets

		// fmt.Printf(" The whole array: %v\n", bookings)
		// fmt.Printf("the type of the array is: %T, and the length is: %v\n", bookings, len(bookings))
		// fmt.Printf("welcome %v, your email is %v\n", userName, userEmail)

		// fmt.Printf("%v", remainingTickets)
		userFirstNames := []string{}
		// for index, elemet := range bookings {
		// 	fmt.Printf("index of %v is the element of: %v", index, elemet)
		// }

		for _, elemet := range bookings {
			firstName := strings.Fields(elemet)[0]
			userFirstNames = append(userFirstNames, firstName)
		}

		if remainingTickets == 0 {
			fmt.Println("Our conferece is booked out")
			break
		}
		//
	}

	// for i := 0; i < len(bookings); i++ {
	// 	name := strings.Fields(bookings[i])[0]
	// 	userFirstNames = append(userFirstNames, name)
	// }

}
