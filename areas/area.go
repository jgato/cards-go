package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func main() {
	sq := square{
		side: 3.3,
	}
	tr := triangle{1.1, 4.3}
	printArea(sq)
	printArea(tr)
}

func printArea(sp shape) {
	fmt.Println("The area of the shape: ", sp.getArea())
}
func (sq square) getArea() float64 {
	return sq.side * sq.side
}

func (tr triangle) getArea() float64 {
	return (tr.base * tr.height) / 2
}
