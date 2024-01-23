package main

import "fmt"

func main() {
	mp := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	var ip int
	fmt.Println("Enter Number : ")
	fmt.Scanln(&ip)

	/*
		if ip < 1 || ip > 7 {
			fmt.Println("Not day")
		} else {
			fmt.Println("Day: ", mp[ip])
		}
	*/

	val, flag := mp[ip]
	if flag {
		fmt.Println(val)
	} else {
		fmt.Println("Not day")
	}

}
