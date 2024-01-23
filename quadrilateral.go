package main

import "fmt"

type quadrilateral interface {
	Area() int
	perimeter() int
}

type rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (dimension rectangle) Area() int {
	// fmt.Print("Area",dimension.length*dimension.breadth)
	return dimension.length * dimension.breadth
}

func (dimension square) Area() int {
	return dimension.side * dimension.side
}

func (dimension rectangle) perimeter() int {
	return 2 * (dimension.length + dimension.breadth)
}
func (dimension square) perimeter() int {
	return 4 * dimension.side
}

func PrintValue(q quadrilateral) {
	fmt.Println("The area :  ", q.Area())
	fmt.Println("The Perimeter : ", q.perimeter())
}

func rectangleProperty() {
	var l, b int
	fmt.Println("Enter length: ")
	fmt.Scan(&l)

	fmt.Println("Enter breadth: ")
	fmt.Scan(&b)

	var rec quadrilateral
	rec = rectangle{l, b}
	PrintValue(rec)

}

func squareProperty() {
	var side int
	fmt.Println("Enter side: ")
	fmt.Scan(&side)

	var sq quadrilateral
	sq = square{side}
	PrintValue(sq)

}
func main() {
	fmt.Println("1.Rectangle\n2.Square\nEnter choice:")
	var ch int
	fmt.Scan(&ch)
	switch ch {
	case 1:
		rectangleProperty()
	case 2:
		squareProperty()
	default:
		fmt.Println("Enter valid choice")

	}
}
