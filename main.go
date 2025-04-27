package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// takes in user input to create bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func main() {
	mybill := createBill()

	fmt.Println(mybill)
}
