package main

import (
	"fmt"
	//this is the import part of the input pattern
	"github.com/gophercoders/simpleio"
)

func main() {
	// This is the variable declaration for the temperature that you type in.
	var temperature int

	// this is the varaible decleration for the temperature in London today
	var temperatureInLondon int
	var temperatureInDublin int
	var temperatureInFrankfurt int
	var temperatureInHongKong int
	var temperatureInMexicoCity int
	var temperatureInMontreal int

	// You need to add 5 more variabels here. One for each city that you picked

	temperatureInLondon = 11 // this needs to be updated on the day!
	temperatureInDublin = 12
	temperatureInFrankfurt = 16
	temperatureInHongKong = 31
	temperatureInMexicoCity = 18
	temperatureInMontreal = 22

	// tell the user what the program does
	fmt.Println("The worldtemperature program tells you which cities are ")
	fmt.Println("hotter or colder than where you live.")

	fmt.Println("Enter the temperature in degrees Celsius today: ")
	// This is the using part of the input pattern.
	// Here you are assigning the number that the user types in to the
	// temperature variable.
	temperature = simpleio.ReadNumberFromKeyboard()

	// This is the selection pattern
	// If the temperature you typed in is less than the temperature in London
	// then the program will print out
	//
	// Colder than London, which is 11 degrees Celsius today.
	//
	// otherwise this if test is skippe dover and nothing is printed out.
	if temperature < temperatureInLondon {
		fmt.Print("Colder than London, which is ")
		fmt.Print(temperatureInLondon)
		fmt.Println(" degrees Celsius today.")
	}
	// if the temperature you typed in is the same as the temperature in London
	// then the program will print out
	//
	// Exactly the same as London.
	//
	if temperature == temperatureInLondon {
		fmt.Println("Exactly the same as London.")
	}
	// Just like the first if test, except this one only works if the
	// temperature you typed in is greater than the temperature in London.
	if temperature > temperatureInLondon {
		fmt.Print("Hotter than London, which is ")
		fmt.Print(temperatureInLondon)
		fmt.Println(" degrees Celsius today.")
	}
	// Add the if statements for the 5 cities that you picked here
	// Dublin
	if temperature < temperatureInDublin {
		fmt.Print("Colder than Dublin, which is ")
		fmt.Print(temperatureInDublin)
		fmt.Println(" degrees Celsius today.")
	}
	if temperature == temperatureInDublin {
		fmt.Println("Exactly the same as Dublin.")
	}
	if temperature > temperatureInDublin {
		fmt.Print("Hotter than Dublin, which is ")
		fmt.Print(temperatureInDublin)
		fmt.Println(" degrees Celsius today.")
	}
	// Frankfurt
	if temperature < temperatureInFrankfurt {
		fmt.Print("Colder than Frankfurt, which is ")
		fmt.Print(temperatureInFrankfurt)
		fmt.Println(" degrees Celsius today.")
	}
	if temperature == temperatureInFrankfurt {
		fmt.Println("Exactly the same as Frankfurt.")
	}
	if temperature > temperatureInFrankfurt {
		fmt.Print("Hotter than Frankfurt, which is ")
		fmt.Print(temperatureInFrankfurt)
		fmt.Println(" degrees Celsius today.")
	}
	// Hong Kong
	if temperature < temperatureInHongKong {
		fmt.Print("Colder than Hong Kong, which is ")
		fmt.Print(temperatureInHongKong)
		fmt.Println(" degrees Celsius today.")
	}
	if temperature == temperatureInHongKong {
		fmt.Println("Exactly the same as Hong Kong.")
	}
	if temperature > temperatureInHongKong {
		fmt.Print("Hotter than Hong Kong, which is ")
		fmt.Print(temperatureInFrankfurt)
		fmt.Println(" degrees Celsius today.")
	}
	// Mexico City
	if temperature < temperatureInMexicoCity {
		fmt.Print("Colder than Mexico City, which is ")
		fmt.Print(temperatureInMexicoCity)
		fmt.Println(" degrees Celsius today.")
	}
	if temperature == temperatureInMexicoCity {
		fmt.Println("Exactly the same as MexicoCity.")
	}
	if temperature > temperatureInMexicoCity {
		fmt.Print("Hotter than MexicoCity, which is ")
		fmt.Print(temperatureInMexicoCity)
		fmt.Println(" degrees Celsius today.")
	}
	// Montreal
	if temperature < temperatureInMontreal {
		fmt.Print("Colder than Montreal, which is ")
		fmt.Print(temperatureInMontreal)
		fmt.Println(" degrees Celsius today.")
	}
	if temperature == temperatureInMontreal {
		fmt.Println("Exactly the same as Montreal.")
	}
	if temperature > temperatureInMontreal {
		fmt.Print("Hotter than Montreal, which is ")
		fmt.Print(temperatureInFrankfurt)
		fmt.Println(" degrees Celsius today.")
	}

}
