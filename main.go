package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// takes in user input to create bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create new bill - ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add an item, s - save the bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		} else {
			b.addItem(name, p)
			fmt.Println("Item added - ", name, price)
			promptOptions(b)
		}
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		} else {
			b.updateTip(t)
			fmt.Println("Tip added - ", tip)
			promptOptions(b)
		}
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
	default:
		fmt.Println("Invalid choice. Please provide a valid input.")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(&mybill)
}
