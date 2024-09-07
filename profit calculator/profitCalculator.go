package main

import (
	"fmt"
)

func main(){
	var revenue, expenses, taxRate, resultTax float64

	fmt.Print("Enter the expenses")
	fmt.Scan(&expenses)

	fmt.Print("Whats your revenue")
	fmt.Scan(&revenue)

	fmt.Print("How much for the task rate?")
	fmt.Scan(&taxRate) //this ampersn acts like a pointer 

	
	resultTax = revenue * (taxRate/ 100)
	resultTax = revenue - resultTax

	//lets take a look at Printf
	fmt.Printf("Revenue is: %.2f", revenue) //with printf you can pass the values inside the function and the % is the value, %f, %v, %t 
	//the . something is how many decimals you want! to show
	fmt.Println("The expenses were a total of ", expenses) //with println you can pass arguments and separate with comas
	fmt.Println("Your profit before tax is: ", revenue)
	fmt.Println("And your final profit after tax is: ", resultTax)


}