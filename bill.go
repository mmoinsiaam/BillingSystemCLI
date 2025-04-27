package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// makes new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 2.99, "cake": 4.99},
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

	fs += fmt.Sprintf("\n%-25v ...$%0.2f", "Total:", total)
	return fs
}
