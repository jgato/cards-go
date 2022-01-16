package main

import "fmt"

type shape interface {
	getArea() float64
}

type triange struct {
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
	tr := triange{1.1, 4.3}
	printArea(sq)
	printArea(tr)

}

func printArea(sp shape) {
	fmt.Println("The area of the shape: ", sp.getArea())
}
func (sq square) getArea() float64 {
	return sq.side * sq.side
}

func (tr triange) getArea() float64 {
	return (tr.base * tr.height) / 2
}
