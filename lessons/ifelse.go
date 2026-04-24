package lessons

import "fmt"

func Ifelse() {
	var score = 14

	fmt.Println(score)

	if score < 16 || score > 6 {
		fmt.Println("yes")
	}
}
