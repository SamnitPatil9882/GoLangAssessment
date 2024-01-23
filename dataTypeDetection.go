package main

import "fmt"

type datatypeInterface interface {
}

func AcceptAnything(d datatypeInterface) {
	fmt.Printf("This is a value of type %T, %v\n", d, d)
}

type Hello struct {
	demo   int
	sample bool
}

func main() {
	fmt.Println("1.Integer\n2.Booleann\n3.String\n4.Struct\nEnter choice:")
	var ch int
	fmt.Scanln(&ch)
	switch ch {
	case 1:
		var intergValue int
		intergValue = 15
		AcceptAnything(intergValue)
	case 2:
		var boolValue bool
		boolValue = true
		AcceptAnything(boolValue)
	case 3:
		var stringValue string
		stringValue = "Hello World"
		AcceptAnything(stringValue)
	case 4:
		structValue := Hello{25, true}
		AcceptAnything(structValue)
	default:
		fmt.Println("Enter valid choice")

	}
}
