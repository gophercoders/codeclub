package main

import (
	"fmt"
	// this is the import pattern for random numbers
	"github.com/gophercoders/random"
	// Now you need to use the import pattern for input
	"github.com/gophercoders/simpleio"
)

func main() {
	// declare some varaibels.
	// a and b are the variabels we are going to use to hold the question in
	var a int
	var b int
	// answer is the variable we are going to use for your answer
	var answer int

	// Tell the user what the program does.
	fmt.Println("The timesquiz shows you how to use loops.")
	fmt.Println("Can you remember your times tables?")
	fmt.Println("")

	// use random numbers again!
	// The program will pick a number between 1 and 12, including 1 and 12
	// and assign it to the variables a and b
	// If you run the program again, the value of a will be different because
	// the computer will have picked a different number.
	a = random.GetRandomNumberInRange(1, 12)
	b = random.GetRandomNumberInRange(1, 12)

	// Now we want to print the question
	// What is a * b?
	// but with the values of a and b
	fmt.Print("What is ")
	fmt.Print(a)
	fmt.Print(" * ")
	fmt.Print(b)
	fmt.Println("?")

	// now you have to read the users answer from the keyboard
	// using the input using pattern
	answer = simpleio.ReadNumberFromKeyboard()

	// now you need a loop!
	// You have to keep looping around while the users answer is
	// not equal to a * b
	// You need to use the loop pattern for this
	for answer != a*b {
		// your loop statement block starts here!
		// Now you need to use an if statement
		// if the users answer was to large print out
		// Sorry, your gueds was to big
		// otherwise the users guess is too small so print out
		// Sorry your guess was to small
		if answer > a*b {
			fmt.Println("Sorry, your guess was to big.")
		} else {
			fmt.Println("Sorry your guess was to small.")
		}
		// now tell the user to try again
		fmt.Println("Try again ")
		// now you have to use the input pattern again to assign
		// the users next guess to the answer variable
		answer = simpleio.ReadNumberFromKeyboard()
	}
	// Now we are outside the loop! So the user must have guessed
	// correctly. Now you need to print Congratulations.

	fmt.Println("Congratulations! You are correct.")
}
