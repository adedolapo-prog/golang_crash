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

	fmt.Printf("The name %s is such a nice name. I hope he is younger than %d \n", middleName, age)

	// arrays and slices
	// arrays have definite number of values, slices dont
	var arr1 = [3]string{"John", "Val", "Christmas"}
	fmt.Println(arr1)

	//to get the length, you use the len() func and pass the arr as a param
	fmt.Println(len(arr1))

	arr2 := []int{3, 4, 7}

	fmt.Println(arr2)
}
