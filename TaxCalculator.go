package main

import (
	"fmt"
)

func calculateTax(earnings float64) float64 {
	taxRate := 0.15
	ceiling := 30000.0
	var tax float64

	if earnings <= ceiling {
		tax = earnings * taxRate
	} else {
		tax = ceiling*taxRate + (earnings-ceiling)*0.20
	}

	return tax
}

func main() {
	var numCitizens int
	fmt.Print("Enter the number of citizens: ")
	fmt.Scanln(&numCitizens)

	for i := 1; i <= numCitizens; i++ {
		var name string
		var earnings float64

		fmt.Printf("Enter citizen %d's name: ", i)
		fmt.Scanln(&name)

		fmt.Printf("Enter citizen %d's earnings: ", i)
		fmt.Scanln(&earnings)

		tax := calculateTax(earnings)
		fmt.Printf("Total tax for citizen %s is: $%.2f\n", name, tax)
	}
}
