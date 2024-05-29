package main

import (
	"fmt"
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
	const conferaneTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking Application\n", conferanceName)
	fmt.Printf("we have total of %v tickets and %v are still remaining to take\n", conferaneTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n======================\n")

	// printing the TYPE of an VARIALBLE
	// fmt.Printf("%T\n", conferanceName)

	var bookings [50]string

	fmt.Println(bookings)
	var userName string
	// var userTicket int
	var userEmail string

	userName = userInput("full name", userName)
	userEmail = userInput("email address", userEmail)
	// userEmail= userInput[int]("ticket numbers", userTicket)
	bookings[0] = userName
	fmt.Printf(" The whole array: %v\n", bookings)
	fmt.Printf("the type of the array is: %T, and the length is: %v\n", bookings, len(bookings))
	fmt.Printf("welcome %v, your email is %v", userName, userEmail)
	// fmt.Printf("user name %v, userTickets %v", userName, userTicket)

}
