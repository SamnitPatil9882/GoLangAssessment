package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	// var ind1, ind2 int
	scanner := bufio.NewScanner(os.Stdin)
	var ipstr string
	fmt.Println("Enter indexes: ")
	scanner.Scan()
	ipstr = scanner.Text()
	inVal := strings.Split(ipstr, " ")
	if len(inVal) != 2 {
		fmt.Print("Invalid Input")
	} else {
		ind1, err1 := strconv.Atoi(inVal[0])
		ind2, err2 := strconv.Atoi(inVal[1])

		if err1 != nil || err2 != nil {
			log.Fatal("Invalid input. Please enter valid numerical values.")
		} else if ind1 < 0 || ind1 > 7 || ind2 < 0 || ind2 > 7 {
			fmt.Println("Incorrect Indexes")
		} else {
			fmt.Println(arr[:ind1+1])
			fmt.Println(arr[ind1 : ind2+1])
			fmt.Println(arr[ind2:])
		}
	}

}
