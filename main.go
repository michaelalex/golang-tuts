package main

import (
	"fmt"
	variables "golang-tuts/golang-tuts/01-variables"
	loops "golang-tuts/golang-tuts/02-loops"
	functions "golang-tuts/golang-tuts/03-functions"
	conditions "golang-tuts/golang-tuts/04-conditions"
	pointers "golang-tuts/golang-tuts/06-pointers"
)

func variablesLesson() {
	fmt.Println("=== 01 - VARIABLES ===")
	fmt.Println(variables.GetTruthy())
	fmt.Println(variables.GetCoupleOfNumbers())
	fmt.Println(variables.GetHelloWorld())
	fmt.Println()
}

func loopsLesson() {
	fmt.Println("=== 02 - LOOPS ===")
	callback := func(currentIndex int) {
		fmt.Println("Current Index", currentIndex)
	}
	loops.BasicIterate(5, callback)
	loops.Iterate(5, callback)
	fmt.Println()
}

func functionsLesson() {
	fmt.Println("=== 03 - FUNCTIONS ===")
	closureValue := functions.AppendString("Hello ")
	fmt.Println(closureValue("World "))
	fmt.Println(closureValue("How are you?"))
	fmt.Println()
}

func conditionsLesson() {
	fmt.Println("=== 04 - CONDITIONS ===")

	oddCallback := func(value int) {
		fmt.Println("Odd ", value)
	}
	evenCallback := func(value int) {
		fmt.Println("Even ", value)
	}
	conditions.OddEvenCallback(1, oddCallback, evenCallback)
	conditions.OddEvenCallback(2, oddCallback, evenCallback)

	fmt.Printf("Number as string %s", conditions.GetNumberAsString(1))
	fmt.Println()
	fmt.Printf("True: What am I %s", conditions.WhatAmI(true))
	fmt.Println()
	fmt.Printf("1: What am I %s", conditions.WhatAmI(1))
	fmt.Println()
	fmt.Printf("\"value\": What am I %s", conditions.WhatAmI("value"))
	fmt.Println()
}

func pointersLesson() {
	fmt.Println()
	fmt.Println("=== 06 - POINTERS ===")

	value := 0
	pointers.UpdateValue(&value, 10)
	fmt.Println("Value with pointer: ", value)

	pointers.UpdateValueWithoutPointer(value, 100)
	fmt.Println("Value without pointer: ", value)

	fmt.Println()
}

func main() {
	variablesLesson()
	loopsLesson()
	functionsLesson()
	conditionsLesson()
	pointersLesson()
}
