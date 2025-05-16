package pointers

import "fmt"

func Solution3() {
	x := 10
	y := 69

	swap(&x, &y)

	fmt.Println(x, y)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
