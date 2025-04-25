package main

import "fmt"

func main() {

	names := [4]string{"Abe", "Absol", "Akhenatun", "Aaron"}
	scores := []int{22, 45, 66}
	fmt.Println(scores[1])

	fmt.Println(names, len(names))

	scores[1] = 55
	fmt.Println(scores[1])
	scores = append(scores, 21)
	fmt.Println(scores, len(scores))
	rangeOne := names[1:3]
	fmt.Print(rangeOne)
}
