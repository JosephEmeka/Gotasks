package main

import (
	"fmt"
)

type Customer struct {
	accountNumber    int
	beginningBalance int
	totalCharges     int
	totalCredits     int
	creditLimit      int
}

func main() {

	var accountNumber int
	var beginningBalance int
	var totalCharges int
	var totalCredits int
	var creditLimit int

	fmt.Print("Enter account number : ")
	fmt.Scanln(&accountNumber)

	fmt.Print("Enter beginning balance : ")
	fmt.Scanln(&beginningBalance)

	fmt.Print("Enter total charges: ")
	fmt.Scanln(&totalCharges)

	fmt.Print("Enter total credits: ")
	fmt.Scanln(&totalCredits)

	fmt.Print("Enter credit limit: ")
	fmt.Scanln(&creditLimit)

	var customer1 = Customer{accountNumber, beginningBalance, totalCharges, totalCredits, creditLimit}
	checkCreditLimit(customer1)
}

func checkCreditLimit(customer Customer) {
	newBalance := customer.beginningBalance + customer.totalCharges - customer.totalCredits
	fmt.Printf("New balance for account number %d: %d\n", customer.accountNumber, newBalance)

	if newBalance > customer.creditLimit {
		fmt.Println("Credit limit exceeded")
	}
}
