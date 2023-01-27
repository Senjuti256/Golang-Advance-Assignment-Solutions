package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (sq Square) Area() float64 {
	return sq.Side * sq.Side
}

func (re Rectangle) Area() float64 {
	return re.Height * re.Width
}

func main() {

	var re = Rectangle{
		Height: 4,
		Width:  8,
	}

	fmt.Println("Area of Rectangle is ", re.Area())

	sq := Square{
		Side: 3,
	}

	fmt.Println("Area of Square: ", sq.Area())
}
