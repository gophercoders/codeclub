package main

import (
	"fmt"
)

func main() {
	// Declare a variable called myName.
	var myName string
	// Now add a variable for your friends name here.
	var friendsName string

	// Set the value of myName to your name.
	// I'll use Owen :) You have to use your name.
	// You have to use inverted commas because Owen is a string.
	myName = "Owen"
	// now set the vale of your friends name here.
	// Ben is my friends name
	friendsName = "Ben"

	fmt.Print("Hello ")
	// Print out the value of myName.
	// You do not need inverted commas because you want to use
	// the value of the variable myName
	fmt.Println(myName)
	// Now print "and" here
	fmt.Print("and ")

	// and then print out your friends name on the same line
	// as the “and”.
	fmt.Print(friendsName)
	// add the full stop
	fmt.Println(".")

} // don't forget the last brace at the bottom
