package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
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
	stringConcept()
	var rect = rectangle{4, 5}
	var cir = circle{5}
	fmt.Println(calculateArea(rect))
	fmt.Println(calculateArea(cir))
	pointerConcept()
	goRoutineConcept()
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

func stringConcept() {
	var str = "resume"
	fmt.Println()
	fmt.Printf("%v %T\n", str[0], str[0])
	for i, v := range str {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of the string is %v", len("resume"))
	var runeStr = 'e'
	fmt.Println(runeStr)
	var stringBuilder strings.Builder
	var stringSlice = []string{"a", "b", "c"}
	for i := range stringSlice {
		stringBuilder.WriteString(stringSlice[i])
	}
	var catString = stringBuilder.String()
	fmt.Println(catString)
}

type rectangle struct {
	length int
	width  int
}

type circle struct {
	radius int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (c circle) area() int {
	return 3 * c.radius
}

type shape interface {
	area() int
}

func calculateArea(s shape) int {
	return s.area()
}

func pointerConcept() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)
	p = &i
	*p = 1
	fmt.Printf("The value p points to %v\n", *p)
	fmt.Printf("The value of i is %v\n", i)
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("The main arr address is %p\n", &arr)
	var result = square(&arr)
	fmt.Printf("arr in main:  %v\n", arr)
	fmt.Printf("The result arr address is %p\n", &result)
	fmt.Printf("Result in main: %v\n", result)

}

func square(arr *[5]int) [5]int {
	fmt.Printf("The square arr address is %p\n", arr)
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
	return *arr
}

var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"abc", "def", "ghi", "jkl", "mno"}
var result = []string{}

func goRoutineConcept() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total execute time: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", result)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from the database is: %v\n", dbData[i])
	mutex.Lock()
	result = append(result, dbData[i])
	mutex.Unlock()
	wg.Done()
}
