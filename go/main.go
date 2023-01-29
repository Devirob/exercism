package main

import (
	"diffsquares"
	"fmt"
)

func main() {
	var myTest = diffsquares.SquareOfSum(5)
	var my2Test = diffsquares.SumOfSquares(5)
	var mu3Test = diffsquares.Difference(5)

	fmt.Println(myTest)
	fmt.Println(my2Test)
	fmt.Println(mu3Test)
}
