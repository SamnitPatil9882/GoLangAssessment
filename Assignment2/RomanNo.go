package main

import "fmt"

func main() {
	var romanstr string
	var real int
	real = 0
	fmt.Println("Enter Roman Value: ")
	fmt.Scan(&romanstr)
	var prev int
	prev = -1
	var curr int
	n := len(romanstr)
	for i := n - 1; i >= 0; i-- {
		if romanstr[i] == 'I' {
			curr = 1
		} else if romanstr[i] == 'V' {
			curr = 5
		} else if romanstr[i] == 'X' {
			curr = 10
		} else if romanstr[i] == 'L' {
			curr = 50
		} else if romanstr[i] == 'C' {
			curr = 100
		} else if romanstr[i] == 'D' {
			curr = 500
		} else if romanstr[i] == 'M' {
			curr = 1000
		}
		if prev > curr {
			real = real - curr
		} else {
			real = real + curr
		}
		prev = curr

	}
	fmt.Println("Number in decimal: ", real)
}

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
