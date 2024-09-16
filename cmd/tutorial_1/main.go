package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Go started!")

	var intNum int
	fmt.Println(intNum)

	var floatNum float64
	fmt.Println(floatNum)

	var myString string
	fmt.Println(myString)

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool
	fmt.Println(myBoolean)

	myVar := "test"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const pi float32 = 3.14
	fmt.Println(pi)

	result, remainder, error := intDivision(8, 2)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("The result of integer division is %v with remainder is %v", result, remainder)
	}

	switch {
	case error != nil:
		fmt.Println(error.Error())
	case remainder == 0:
		fmt.Printf("The result if integer division is %v", result)
	default:
		fmt.Printf("The result of integer division is %v with remainder is %v", result, remainder)
	}
}

func intDivision(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("can not divide by zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
