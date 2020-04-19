package main

import (
	"fmt"
	variables "golang-tuts/golang-tuts/01-variables"
	loops "golang-tuts/golang-tuts/02-loops"
	functions "golang-tuts/golang-tuts/03-functions"
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

func main() {
	variablesLesson()
	loopsLesson()
	functionsLesson()
}
