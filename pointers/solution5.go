package pointers

import "fmt"

func Solution5() {
	num1 := -9
	num2 := -3

	fmt.Println(minMax(&num1, &num2))
}

func minMax(num1, num2 *int) (int, int) {
	larger := 0
	smaller := 0

	if *num1 >= *num2 {
		larger = *num1
		smaller = *num2
	} else {
		larger = *num2
		smaller = *num1
	}

	return larger, smaller
}
