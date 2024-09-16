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

	arrayConcept()
	sliceConcept()
	mapConcept()
}

func intDivision(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("can not divide by zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}

func arrayConcept() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println()
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:3])
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}

func sliceConcept() {
	var slice []int32 = []int32{4, 5, 6}
	fmt.Printf("The lenght of slice is %v and the capacity is %v\n", len(slice), cap(slice))
	slice = append(slice, 7)
	fmt.Printf("The lenght of slice is %v and the capacity is %v\n", len(slice), cap(slice))

	anotherSlice := make([]int32, 3, 4)
	fmt.Printf("The lenght of slice is %v and the capacity is %v\n", len(anotherSlice), cap(anotherSlice))
	anotherSlice = append(anotherSlice, 5, 6)
	fmt.Printf("The lenght of slice is %v and the capacity is %v\n", len(anotherSlice), cap(anotherSlice))
}

func mapConcept() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"vikash": 32, "kumar": 30}
	age, ok := myMap2["vikash"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("Not valid name")
	}
	for name, age := range myMap2 {
		fmt.Printf("Name: %v and Age: %v\n", name, age)
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}
