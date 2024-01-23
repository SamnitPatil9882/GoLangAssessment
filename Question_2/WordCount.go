package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var ipstr string
	fmt.Println("Enter String: ")
	scanner.Scan()
	ipstr = scanner.Text()
	// fmt.Println(ipstr)

	wordArray := strings.Split(ipstr, " ")
	// fmt.Println(wordArray[0])

	mp := make(map[string]int, 0)
	// var lenIpstr int
	lenIpstr := len(wordArray)
	// var maxFreq int
	maxFreq := 0

	for i := 0; i < lenIpstr; i++ {
		mp[wordArray[i]]++
		if mp[wordArray[i]] > maxFreq {
			maxFreq = mp[wordArray[i]]
		}
	}

	ans := make([]string, 0)

	for i := 0; i < lenIpstr; i++ {
		if mp[wordArray[i]] == maxFreq {
			ans = append(ans, wordArray[i])
			mp[wordArray[i]] = -1
		}
	}
	fmt.Print(ans)

}
