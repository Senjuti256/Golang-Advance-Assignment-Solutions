package main

import "fmt"

func main() {

	var arr = [5]int{45, 46, 67, 88, 99}

	var a [5]int
	a[0] = 45
	a[1] = 46
	a[2] = 67
	a[3] = 88
	a[4] = 99

	sum := arr[2] + arr[4]

	fmt.Println("The sum of elements present in index 2 and 4 is", sum)

	//creating slice
	var b []int = a[3:5]
	fmt.Println(b)

}
