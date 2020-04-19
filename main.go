package main

import (
	"fmt"
	variables "golang-tuts/golang-tuts/01-variables"
	loops "golang-tuts/golang-tuts/02-loops"
)

func main() {
	// 01 - Variables
	fmt.Println(variables.GetTruthy())
	fmt.Println(variables.GetCoupleOfNumbers())
	fmt.Println(variables.GetHelloWorld())

	// 02 - loops
	callback := func(currentIndex int) {
		fmt.Println("Current Index", currentIndex)
	}
	loops.BasicIterate(5, callback)
	loops.Iterate(5, callback)
}
