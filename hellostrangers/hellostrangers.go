package main

import "fmt"

// you need to add the import part here like this

func main() {
	// Declare a variable called myName.
	var myName string
	// Declare a variable called myAge
	var myAge int
	// Now add a variable for your friends name here.
	var friendsName string
	// Now add a variable for your friends age here.
	var friendsAge int

	fmt.Print("Please tell me your name: ")
	// now you have to read a value for your name from the keyboard
	// and assign it to the myName variable
	// You need to use the import pattern for this

	fmt.Print("Please tell me how old you are: ")
	// now you have to read a value for your age from the keyboard
	// and assign it to the myAge variable

	// now you have to read your friends name and age from the keyboard
	// you will need extra print statements to tell the user when to enter
	// your friends name and age

	fmt.Print("Hello ")
	// Print out the value of myName.
	// You do not need inverted commas because you want to use
	// the value of the variable myName
	fmt.Println(myName)
	fmt.Print("You are ")
	// now print out your age
	fmt.Print(myAge)
	fmt.Println(" years old.")
	// Now print out your friends name and age here
	fmt.Print("Hello ")
	fmt.Println(friendsName)
	fmt.Print("You are ")
	fmt.Print(friendsAge)
	fmt.Println(" years old.")

	fmt.Print("Our ages multiplied together is ")
	// now multiply your ages together and print out the answer
	// Use a new veriable for this. It is ok to decalre a new variable here.
	// It doesn't have to be at the top, so long as it is decalred before
	// we first use the variable.
	var ages int
	ages = myAge * friendsAge
	// now print the answer
	fmt.Println(ages)

	fmt.Print("The difference between our ages is ")
	// now print out the difference between your ages
	// We don't have to use a new variable do do this. We can work the anser
	// as we go. Like this, which prints the answer of the sum.
	fmt.Println(myAge - friendsAge)

	fmt.Print("The sum of our ages is ")
	// now print out the sum of your ages
	// We'll use the quick way again
	fmt.Println(myAge + friendsAge)
}
