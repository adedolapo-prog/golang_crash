package main

import "fmt"

func main() {
	// strings
	var myName string = "Huey"
	middleName := "Riley"

	fmt.Println(myName + " " + middleName)

	// numbers and floats
	var age int = 39
	sisterAge := 23

	fmt.Println(fmt.Sprintf("He is %d and his sister is %d", age, sisterAge))

	litresOfFuel := 24.6
	fmt.Println(fmt.Sprintf("%s had to purchase %.2f litres of fuel", middleName, litresOfFuel))
}
