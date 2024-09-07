package main

import (
	"fmt"
	"math"
)

/*
import "fmt"  so instead of going one by one you can just do an import() and extract what you need
import "math"*/

func main() {
	const inflationRate = 2.5
	 var investmentAmount float64  // you can initialize in 0 or just leave the variable ready without initialization by your side
	 expectedReturnRate := 0.0 //you can write the type of number just here!, so you save some lines of code and dont convert
	 years := 0.0
	//var singleChar rune = 'a'

	fmt.Print("Investment Amount> ")
	fmt.Scan(&investmentAmount)
	
	fmt.Print("Years> ")
	fmt.Scan(&years)

	fmt.Print("What is your expected return rate?> ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	/*and just like in C++ if you want to make any operations with distinct type of values
	you must convert one value to the other to match types*/
	futureValue = math.Floor(futureValue)
	futureRealValue := futureValue/ math.Pow(1+inflationRate/100, years)

	fmt.Println("and the value is: ", futureValue)
	fmt.Println("this is the real one: ", futureRealValue)
}
