package main

import "fmt"

func main() {

	f := 67.88
	var in int
	in = int(f)
	fmt.Println(in)

	var num1 int
	var num2 float64
	var sum float64
	//fmt.Println("Enter the num1")
	//fmt.Scanln(&num1)
	//fmt.Println("Enter the num2")
	//fmt.Scanln(&num2)

	num1 = 8
	num2 = 8.6
	sum = float64(num1) + num2
	fmt.Println(sum)
}
