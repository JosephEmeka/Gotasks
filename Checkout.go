package main

import (
	"fmt"
)

type Item struct {
	ItemName    string
	Price       float64
	Quantity    int
	TotalAmount float64
}

type CheckOutApp struct {
	ItemList        []Item
	CustomerName    string
	CashierName     string
	Discount        float64
	VATPercentage   float64
	MyCalender      string
	CustomerPayment float64
	CustomerBalance float64
}

func (app *CheckOutApp) AddPurchasedItem(itemBought string, price float64, quantity int) {
	totalAmount := price * float64(quantity)
	item := Item{ItemName: itemBought, Price: price, Quantity: quantity, TotalAmount: totalAmount}
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

func (app *CheckOutApp) CashierNameMenu() {
	fmt.Println("What is your name? ")
	fmt.Scanln(&app.CashierName)

	fmt.Println("How much discount will he get? ")
	fmt.Scanln(&app.Discount)
}

func (app *CheckOutApp) PrintItemDetails() {
	fmt.Println("  		       ")
	fmt.Println("                   ")
	fmt.Println("                   ")
	fmt.Println("  SEMICOLON  STORES")
	fmt.Println("  MAIN BRANCH")
	fmt.Println("  LOCATION: 312, HERBERT MACAULAY WAY, SABO YABA, LAGOS.")
	fmt.Println("  TELL:  03293828343")
	fmt.Println("  DATE: ", app.MyCalender)
	fmt.Println("  Cashier: ", app.CashierName)
	fmt.Println("  Customer Name: ", app.CustomerName)
	fmt.Println("                   ")
	fmt.Println("==================================================================")
	fmt.Println("  ITEM \t\tPRICE \t\tQTY \t\tTOTAL(NGN) ")
	fmt.Println("-------------------------------------------------------------------")
	for _, item := range app.ItemList {
		fmt.Printf("  %s \t\t%.2f \t\t%d \t\t%.2f\n", item.ItemName, item.Price, item.Quantity, item.TotalAmount)
	}
	fmt.Println("-------------------------------------------------------------------")
	fmt.Printf("\t\t\t\tSub Total:        %.2f\n", app.SubTotal())
	fmt.Printf("\t\t\t\tDiscount:         %.2f\n", app.calculateDiscount())
	fmt.Printf("\t\t\t\tVAT @ 17.50%% :    %.2f\n", app.calculateVAT())
	fmt.Println("==================================================================")
	fmt.Printf("\t\t\t\tBill Total:        %.2f\n", app.totalBill())
	fmt.Println("==================================================================")
	fmt.Println("THIS IS NOT A RECEIPT KINDLY PAY ", app.totalBill())
	fmt.Println("==================================================================")
}

func (app *CheckOutApp) AmountPaid() float64 {
	app.CustomerBalance = app.CustomerPayment
	return app.CustomerBalance
}

func (app *CheckOutApp) calculateDiscount() float64 {
	return (app.Discount / 100) * app.SubTotal()
}

func (app *CheckOutApp) calculateVAT() float64 {
	total := app.SubTotal()
	return (app.VATPercentage / 100) * total
}

func (app *CheckOutApp) totalBill() float64 {
	subTotal := app.SubTotal()
	finalDiscount := app.calculateDiscount()
	vat := app.calculateVAT()
	return subTotal + finalDiscount + vat
}

func (app *CheckOutApp) CustomersInvoice() {
	app.PrintItemDetails()
}

func main() {
	app := CheckOutApp{}
	app.StartMenu()
	app.CashierNameMenu()
	app.CustomersInvoice()
}
