package main

import "fmt"

func main() {

	names := [4]string{"Abe", "Absol", "Akhenatun", "Aaron"}
	for index, value := range names {
		fmt.Println("Pos is:", index, "value is:", value)
	}
}
