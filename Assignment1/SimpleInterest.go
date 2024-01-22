package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
)

func calculateSimpleInterest(principal, rate, time float64) float64 {
	return principal * rate * time / 100
}

func main() {
	var userInput string

	fmt.Print("Enter principal, rate, time (comma-separated): ")
	fmt.Scanln(&userInput)

	
	inputValues := strings.Split(userInput, ",")

	if len(inputValues) != 3 {
		log.Fatal("Invalid input format. Please provide principal, rate, and time.")
	}

	
	principal, err1 := strconv.ParseFloat(strings.TrimSpace(inputValues[0]), 64)
	rate, err2 := strconv.ParseFloat(strings.TrimSpace(inputValues[1]), 64)
	time, err3 := strconv.ParseFloat(strings.TrimSpace(inputValues[2]), 64)

	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatal("Invalid input. Please enter valid numerical values.")
	}

	
	interest := calculateSimpleInterest(principal, rate, time)

	
	fmt.Printf("Simple Interest: %.2f\n", interest)
}

