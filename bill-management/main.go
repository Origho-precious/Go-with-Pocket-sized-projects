package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getInput(prompt string, reader *bufio.Reader) (string, error){
	fmt.Print(prompt)

	response, err := reader.ReadString('\n')

	response = strings.TrimSpace(response)

	return response, err
}

func createNewBill() bill {
	reader := bufio.NewReader(os.Stdin)

	billName, _ := getInput("What do you want to call your new bill? ", reader)

	todaysBill := newBill(billName)

	fmt.Println("Bill created successfully -", billName)

	return todaysBill
}

func handleItemAddition(b *bill, reader * bufio.Reader) {
	itemName, _ := getInput("Item name: ", reader)
	price, _ := getInput("Price ($) e.g 4.99: ", reader)

	floatPrice, err := strconv.ParseFloat(price, 64)
		
	if err != nil {
		fmt.Println("Price must be a number e.g 5.00, 4.99 etc")
		handleUserActions(b)
		return 
	}

	b.addItem(itemName, floatPrice)


	fmt.Println("Item added successfully!", itemName, ":", price)
	handleUserActions(b)
}

func saveBill(b *bill) {
	fmt.Println("You chose to save your bill(s)")
	b.saveBill()
}

func handleTipAddition(b *bill, reader * bufio.Reader) {
	price, _ := getInput("Tip amount ($) e.g 4.55: ", reader)

	floatPrice, err := strconv.ParseFloat(price, 64)

	if err != nil {
		fmt.Println("Tip amount must be a number e.g 5.00, 4.99 etc")
		handleUserActions(b)
		return 
	}

	b.addTip(floatPrice)

	fmt.Printf("Added $%v tip to the bill!\n", price)
	handleUserActions(b)
}

func handleUserActions(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	action, _ := getInput(
		"Choose option (a - add item, s - save bill, t - add tip) ", reader,
	)

	switch action {
		case "a": 
			handleItemAddition(b, reader)
		case "s":
			saveBill(b)
		case "t":
			handleTipAddition(b, reader)
		default:
			fmt.Println("Wrong option!", action)
			handleUserActions(b)
	}
}

func main () {
	todaysBill := createNewBill()
	handleUserActions(&todaysBill)
}