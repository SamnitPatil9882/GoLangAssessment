package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func parse(alice chan string, bob chan string, sgnl chan bool, str string) {
	n := len(str)
	temp := ""
	for i := 0; i < n; i++ {
		ch := str[i]
		switch ch {
		case '$':
			alice <- temp
			temp = ""
			// ch=''
		case '#':
			bob <- temp
			temp = ""
		case '^':
			// close(alice)
			// close(bob)
			sgnl <- true
			wg.Done()
			return
		default:
			temp += string(ch)

		}
	}
}
func printres(alice chan string, bob chan string, sgnl chan bool) {
	//defer wg.Done()
	for {
		select {
		case <-sgnl:
			close(alice)
			close(bob)
			// close(sgnl)
			wg.Done()
			return
		case res := <-alice:
			fmt.Println("alice:", res)
		case res := <-bob:
			fmt.Println("bob:", res)
		}
	}
}
func main() {
	var str string
	str = "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

	alice := make(chan string)
	bob := make(chan string)
	sgnl := make(chan bool)

	fmt.Println("Enter string: ")
	fmt.Scan(&str)

	wg.Add(2)
	go parse(alice, bob, sgnl, str)
	go printres(alice, bob, sgnl)
	wg.Wait()

}
