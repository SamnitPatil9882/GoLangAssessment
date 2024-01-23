package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rectangle struct {
	length  float64
	breadth float64
}

func area(rec rectangle) {
	fmt.Println("Area of Rectngle:", int(rec.length*rec.breadth))
}
func perimeter(rec rectangle) {
	fmt.Println("Area of Rectngle:", int(2*(rec.length+rec.breadth)))
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ipstr string
	fmt.Println("Enter length and breadth:  ")
	scanner.Scan()
	ipstr = scanner.Text()
	inVal := strings.Split(ipstr, " ")
	if len(inVal) != 2 {
		fmt.Print("Invalid Input")
	} else {
		len, err1 := strconv.ParseFloat(strings.Trim(inVal[0], " "), 64)
		br, err2 := strconv.ParseFloat(strings.Trim(inVal[1], " "), 64)

		if err1 != nil || err2 != nil {
			log.Fatal("Invalid input. Please enter valid numerical values.")
		} else {
			rec := rectangle{len, br}
			area(rec)
			perimeter(rec)
		}
	}

}
