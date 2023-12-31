package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	newlyCreatedBill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.0,
	}

	return newlyCreatedBill
}

func (b *bill) format() string {
	formattedString := "\nBill breakdown: \n"
	var total float64 = 0.0

	for itemName, amount := range b.items {
		formattedString += fmt.Sprintf("%-15v $%v \n", itemName+":", amount)
		total += amount
	}

	formattedString += fmt.Sprintf("%-15v $%0.2f\n\n", "tip:", b.tip)
	total += b.tip

	formattedString += fmt.Sprintf("%-15v $%0.2f \n", "total:", total)

	return formattedString
}

func (b *bill) addTip(amount float64) {
	b.tip = amount
}

func (b *bill) addItem(itemName string, price float64) {
	// fmt.Println("Got here", itemName, price)

	prevPrice, ok := b.items[itemName]

	if !ok {
		b.items[itemName] = price
	} else {
		b.items[itemName] = price + prevPrice
	}

	// fmt.Println("UpdatedItems:", b.items)
}

func (b *bill) saveBill() {
	data := []byte(b.format())

	filePath := fmt.Sprintf("bills/%v.txt", b.name)

	f, openError := os.OpenFile(
		filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)

	if openError != nil {
		panic(openError)
	}

	_, writeError := f.Write(data)

	if writeError != nil {
		f.Close()
		panic(openError)
	}

	fmt.Println("Bill was saved to file:", filePath)
}
