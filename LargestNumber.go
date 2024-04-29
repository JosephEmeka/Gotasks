package main

import (
	"fmt"
)

func main() {
	var counter, number, largest int
	for counter = 0; counter < 10; counter++ {
		fmt.Println("Enter number: ")
		if _, err := fmt.Scanln(&number); err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if number > largest {
			largest = number
		}
	}
	fmt.Printf("This is the largest number: %d", largest)
}
