package main

import "fmt"

func reverseString(str string, ch chan string) {
	temp := ""
	n := len(str)
	for i := n - 1; i >= 0; i-- {
		temp += string(str[i])
	}
	ch <- temp

}
func main() {
	ch := make(chan string)
	var str string
	fmt.Println("Enter string")
	fmt.Scan(&str)

	go reverseString(str, ch)
	revStr := <-ch

	fmt.Println("reversed String: ", revStr)

	// go reverseString()
}
