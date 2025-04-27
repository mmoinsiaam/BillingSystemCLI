package main

import "fmt"

func main() {
	b := newBill("Mario's Bill")
	b.addItem("pie", 5.89)
	b.updateTip(6.77)
	fmt.Println(b.format())
}
