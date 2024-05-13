package main

import (
	"fmt"
	"time"
)

type Item struct {
	ItemName    string
	Price       float64
	Quantity    int
	TotalAmount float64
}

type CheckOutApp struct {
	ItemList        []Item
	MyCalender      time.Time
	CashierName     string
	CustomerName    string
	VATPercentage   float64
	Discount        float64
	CustomerPayment float64
	CustomerBalance float64
}

func (app *CheckOutApp) AddPurchasedItem(itemName string, price float64, quantity int) {
	totalAmount := price * float64(quantity)
	item := Item{ItemName: itemName, Price: price, Quantity: quantity, TotalAmount: totalAmount}
	app.ItemList = append(app.ItemList, item)
}

func (app *CheckOutApp) SubTotal() float64 {
	total := 0.0
	for _, item := range app.ItemList {
		total += item.TotalAmount
	}
	return total
}

func (app *CheckOutApp) StartMenu() {
	fmt.Println("What is the customer's Name? ")
	fmt.Scanln(&app.CustomerName)

	for {
		fmt.Println("What did the user buy? ")
		var itemBought string
		fmt.Scanln(&itemBought)

		fmt.Println("How many pieces? ")
		var quantity int
		fmt.Scanln(&quantity)

		fmt.Println("How much per unit? ")
		var price float64
		fmt.Scanln(&price)

		app.AddPurchasedItem(itemBought, price, quantity)

		fmt.Println("Add more Items? Type Yes or No ")
		var moreItems string
		fmt.Scanln(&moreItems)

		if moreItems == "No" {
			break
		}
	}
}

func main() {
	app := CheckOutApp{}
	app.StartMenu()
}
