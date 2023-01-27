package main

import "fmt"

func main() {
	f := 4.5
	p := &f
	fmt.Printf("The value of f is %f \n", *p)
	fmt.Printf("The address of f is %v \n", p)
	fmt.Println("Value of f before updation", f)

	*p = 9.8

	fmt.Println("Value of f after updation", f)

}
