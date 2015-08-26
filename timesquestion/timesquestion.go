package main

import (
	"fmt"
	// This is new!
	"github.com/gophercoders/random"
	// You need to use the import pattern for input againhere
)

func main() {
	// declare some varaibels.
	// a and b are the variabels we are going to use to hold the question in
	var a int
	var b int
	// answer is the variable we are going to use for your answer
	var answer int

	// Tell the user what the program does.
	fmt.Println("The timesquestion program shows you how to use if and else.")
	fmt.Println("Can you remember your times tables?")
	fmt.Println("")

	// This new!
	// a and b are both set to a random number.
	// The program will pick a number between 1 and 12, including 1 and 12
	// and assign it to the variable a
	// If you run the program again, the value of a will be different because
	// the computer will have picked a different number.
	a = random.GetRandomNumberInRange(1, 12)
	// b is also picked randomly!
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

	// Now you have to fill in the rest!
	// Use the selection if else if pattern...

	// if the users gues is exacly a*b then you have to print
	// congratulations

	// else if the users guess is to large print out
	// Sorry you guess was to big and then print out the correct answer
	// else the users guess must have been too small so print out
	// Sorry your guess was too small and then print out the correct
	// answer

	// this sould be the last line!
	// Run the program again and you'll be asked a different question!
	fmt.Println("Run the program again to try another question.")
}
