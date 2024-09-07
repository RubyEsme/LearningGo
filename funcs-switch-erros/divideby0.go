package main

import (
	"errors" //this is the library that handles errors of functions wrong executed
	"fmt"
)

func main() {
	//create a function that receives two integers and divides them, but handle the errors
	//also create a function that receives a message and prints it
	printMessage("Heyyy, I did this!")
	var resultDivision, reminder, error = divide(5,3)
	switch {
	case error != nil:
		fmt.Println(error.Error())
	case reminder == 0:
		fmt.Printf("The result of the division is %v", resultDivision)
	default:
		fmt.Printf("The result of the division is %v, and the reminder is %v ", resultDivision, reminder)

	}
}

func printMessage(message string) {
	fmt.Println(message)
}
func divide(first int, second int) (int, int, error) {

	var err error //create the variable of type error, add a message using the library errors with the New() and return it

	if second == 0 {
		err = errors.New("Hey, please be kind and give me a nice denomitaror>0")
		return 0, 0, err //returns 0,0, err because of the things that the function is going to return no matter what
	}

	var result int = first / second
	reminder := first % second

	return result, reminder, err
}
