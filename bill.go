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

// makes new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// receiver function to format bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	//list items with price
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%0.2f\n", k+":", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Tip:", b.tip)
	total += b.tip
	//add total
	fs += fmt.Sprintf("\n%-25v ...$%0.2f", "Total:", total)
	return fs
}

// updates tip
func (b *bill) updateTip(tips float64) {
	b.tip = tips
}

// adds item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Bill was saved to file.")
	}
}
